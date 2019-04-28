package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/brotherlogic/godiscogs"
	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/goserver/utils"
	"github.com/brotherlogic/keystore/client"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbd "github.com/brotherlogic/godiscogs"
	pbg "github.com/brotherlogic/goserver/proto"
	pb "github.com/brotherlogic/recordcollection/proto"
	pbrm "github.com/brotherlogic/recordmover/proto"
	pbrp "github.com/brotherlogic/recordprocess/proto"
	pbro "github.com/brotherlogic/recordsorganiser/proto"

	_ "net/http/pprof"
)

type quotaChecker interface {
	hasQuota(ctx context.Context, folder int32) (*pbro.QuotaResponse, error)
}

type moveRecorder interface {
	moveRecord(record *pb.Record, oldFolder, newFolder int32) error
}

type prodMoveRecorder struct{}

func (p *prodMoveRecorder) moveRecord(record *pb.Record, oldFolder, newFolder int32) error {
	ip, port, err := utils.Resolve("recordmover")
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	rmclient := pbrm.NewMoveServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	_, err = rmclient.RecordMove(ctx, &pbrm.MoveRequest{Move: &pbrm.RecordMove{
		InstanceId: record.GetRelease().InstanceId,
		FromFolder: oldFolder,
		ToFolder:   newFolder,
		Record:     record,
	}})

	return err
}

type prodQuotaChecker struct{}

func (p *prodQuotaChecker) hasQuota(ctx context.Context, folder int32) (*pbro.QuotaResponse, error) {
	ip, port, err := utils.Resolve("recordsorganiser")
	if err != nil {
		return &pbro.QuotaResponse{}, err
	}

	conn, err := grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
	if err != nil {
		return &pbro.QuotaResponse{}, err
	}
	defer conn.Close()

	client := pbro.NewOrganiserServiceClient(conn)
	return client.GetQuota(ctx, &pbro.QuotaRequest{FolderId: folder})
}

type saver interface {
	GetCollection() []*godiscogs.Release
	GetWantlist() ([]*godiscogs.Release, error)
	GetRelease(id int32) (*godiscogs.Release, error)
	AddToFolder(folderID int32, releaseID int32) (int, error)
	SetRating(releaseID int, rating int) error
	MoveToFolder(folderID, releaseID, instanceID, newFolderID int) string
	DeleteInstance(folderID, releaseID, instanceID int) string
	SellRecord(releaseID int, price float32, state string) int
	GetSalePrice(releaseID int) float32
	RemoveFromWantlist(releaseID int)
	AddToWantlist(releaseID int)
	UpdateSalePrice(saleID int, releaseID int, condition string, price float32) error
	GetCurrentSalePrice(saleID int) float32
	GetCurrentSaleState(saleID int) godiscogs.SaleState
	RemoveFromSale(saleID int, releaseID int) error
}

type scorer interface {
	GetScore(ctx context.Context, instanceID int32) (float32, error)
}

type prodScorer struct {
	dial func(server string) (*grpc.ClientConn, error)
}

func (p *prodScorer) GetScore(ctx context.Context, instanceID int32) (float32, error) {
	conn, err := p.dial("recordprocess")
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	client := pbrp.NewScoreServiceClient(conn)
	res, err := client.GetScore(ctx, &pbrp.GetScoreRequest{InstanceId: instanceID})
	if err != nil {
		return -1, err
	}

	score := float32(0)
	for _, sc := range res.Scores {
		score += float32(sc.Rating)
	}

	return score / float32(len(res.Scores)), nil
}

//Server main server type
type Server struct {
	*goserver.GoServer
	collection     *pb.RecordCollection
	retr           saver
	lastSyncTime   time.Time
	lastPushTime   time.Time
	lastPushLength time.Duration
	lastPushDone   int
	lastPushSize   int
	cacheWait      time.Duration
	pushMutex      *sync.Mutex
	pushMap        map[int32]*pb.Record
	pushWait       time.Duration
	saveNeeded     bool
	quota          quotaChecker
	mover          moveRecorder
	nextPush       *pb.Record
	lastWantUpdate int32
	wantCheck      string
	lastWantText   string
	scorer         scorer
	saleMap        map[int32]*pb.Record
	lastSalePush   time.Time
	lastSyncLength time.Duration
	salesPushes    int64
	soldAdjust     int64
	wantUpdate     string
	saves          int64
	saveMutex      *sync.Mutex
	biggest        int64
}

func (s *Server) findBiggest(ctx context.Context) error {
	biggestb := 0
	for _, r := range s.collection.GetRecords() {
		data, _ := proto.Marshal(r)
		if len(data) > biggestb {
			s.biggest = int64(r.GetRelease().Id)
			biggestb = len(data)
		}
	}
	return nil
}

const (
	// KEY The main collection
	KEY = "/github.com/brotherlogic/recordcollection/collection"

	//TOKEN for discogs
	TOKEN = "/github.com/brotherlogic/recordcollection/token"
)

func (s *Server) readRecordCollection(ctx context.Context) error {
	collection := &pb.RecordCollection{}
	data, _, err := s.KSclient.Read(ctx, KEY, collection)

	if err != nil {
		return err
	}

	s.collection = data.(*pb.RecordCollection)

	//Fill the push map
	for _, r := range s.collection.GetRecords() {

		//Copy over the instance id if needed
		if r.GetMetadata().InstanceId == 0 {
			r.GetMetadata().InstanceId = r.GetRelease().InstanceId
		}

		// Stop repeated fields from blowing up
		if len(r.GetRelease().GetFormats()) > 100 {
			r.GetRelease().Images = []*pbd.Image{}
			r.GetRelease().Artists = []*pbd.Artist{}
			r.GetRelease().Formats = []*pbd.Format{}
			r.GetRelease().Labels = []*pbd.Label{}
			r.GetMetadata().LastCache = 1
		}

		if r.GetMetadata().GetMoveFolder() > 0 || r.GetMetadata().GetSetRating() > 0 {
			r.GetMetadata().Dirty = true
		}

		if r.GetMetadata().Dirty {
			s.pushMutex.Lock()
			s.pushMap[r.GetRelease().InstanceId] = r
			s.pushMutex.Unlock()
		}

		if len(r.GetRelease().GetTracklist()) == 0 {
			r.GetMetadata().LastCache = 1
		}
	}

	// Fill the sale map
	for _, r := range s.collection.GetRecords() {
		if r.GetMetadata().SaleId > 0 {
			s.saleMap[r.GetMetadata().SaleId] = r
		}
	}

	return nil
}

func (s *Server) saveRecordCollection(ctx context.Context) {
	s.saveMutex.Lock()
	defer s.saveMutex.Unlock()
	s.saves++
	s.KSclient.Save(ctx, KEY, s.collection)
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterRecordCollectionServiceServer(server, s)
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	s.saveRecordCollection(ctx)
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	if master {
		tType := &pb.Token{}
		tResp, _, err := s.KSclient.Read(context.Background(), TOKEN, tType)
		if err != nil {
			return err
		}
		s.retr = pbd.NewDiscogsRetriever(tResp.(*pb.Token).Token, s.Log)

		err = s.readRecordCollection(ctx)

		return err
	}

	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	noGoal := int64(0)
	example := int64(0)
	needsBump := int64(0)
	nbExample := int64(0)
	numUnlistened := int64(0)
	for _, r := range s.collection.GetRecords() {
		if r.GetMetadata().Category == pb.ReleaseMetadata_UNLISTENED {
			numUnlistened++
		}
		if r.GetMetadata().GoalFolder == 0 {
			noGoal++
			example = int64(r.GetRelease().Id)
		}
		if r.GetMetadata().DateAdded > 1555311600 {
			nbExample = int64(r.GetRelease().InstanceId)
			needsBump++
		}
	}
	return []*pbg.State{
		&pbg.State{Key: "no_goal", Value: noGoal},
		&pbg.State{Key: "no_goal_example", Value: example},
		&pbg.State{Key: "needs_date", Value: needsBump},
		&pbg.State{Key: "needs_date_example", Value: nbExample},
		&pbg.State{Key: "unlistened", Value: numUnlistened},
	}
}

// Init builds out a server
func Init() *Server {
	s := &Server{
		GoServer:       &goserver.GoServer{},
		lastSyncTime:   time.Now(),
		pushMap:        make(map[int32]*pb.Record),
		pushWait:       time.Minute,
		pushMutex:      &sync.Mutex{},
		lastPushTime:   time.Now(),
		lastPushSize:   0,
		lastPushLength: 0,
		quota:          &prodQuotaChecker{},
		mover:          &prodMoveRecorder{},
		lastWantText:   "",
		scorer:         &prodScorer{},
		saleMap:        make(map[int32]*pb.Record),
		lastSalePush:   time.Now(),
		wantUpdate:     "unknown",
		saveMutex:      &sync.Mutex{},
	}
	s.scorer = &prodScorer{s.DialMaster}
	return s
}

func (s *Server) cacheLoop(ctx context.Context) error {
	for _, r := range s.collection.GetRecords() {
		if r.GetMetadata().LastCache <= 1 {
			s.cacheRecord(ctx, r)
			return nil
		}
	}
	return nil
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	var token = flag.String("token", "", "Discogs token")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.GoServer.KSclient = *keystoreclient.GetClient(server.GetIP)

	server.PrepServer()
	if len(*token) > 0 {
		server.KSclient.Save(context.Background(), TOKEN, &pb.Token{Token: *token})
		log.Fatalf("Written TOKEN")
	}

	server.Register = server
	server.SendTrace = true

	t := time.Now()
	err := server.RegisterServer("recordcollection", false)
	if err != nil {
		log.Fatalf("Unable to register (%v) at %v", err, time.Now().Sub(t))
	}

	// This enables pprof
	go http.ListenAndServe(":8089", nil)

	server.RegisterRepeatingTask(server.runSync, "run_sync", time.Hour)
	server.RegisterRepeatingTask(server.runSyncWants, "run_sync_wants", time.Hour)
	server.RegisterRepeatingTask(server.pushWants, "push_wants", time.Minute)
	server.RegisterRepeatingTask(server.runPush, "run_push", time.Minute)
	server.RegisterRepeatingTask(server.syncIssue, "sync_issue", time.Hour)
	server.RegisterRepeatingTask(server.pushSales, "push_sales", time.Minute)
	server.RegisterRepeatingTask(server.cacheLoop, "cache_loop", time.Minute)
	server.RegisterRepeatingTask(server.findBiggest, "find_biggest", time.Minute*5)

	server.Serve()
}
