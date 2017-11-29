package main

import (
	"flag"
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/brotherlogic/godiscogs"
	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"

	pbg "github.com/brotherlogic/goserver/proto"
	pb "github.com/brotherlogic/recordcollection/proto"
)

type saver interface {
	GetCollection() []godiscogs.Release
	GetWantlist() []godiscogs.Release
}

//Server main server type
type Server struct {
	*goserver.GoServer
	collection   *pb.RecordCollection
	retr         saver
	lastSyncTime time.Time
}

const (
	// KEY The main collection
	KEY = "github.com/brotherlogic/recordcollection/collection"
)

func (s *Server) readRecordCollection() error {
	collection := &pb.RecordCollection{}
	data, _, err := s.KSclient.Read(KEY, collection)

	if err != nil {
		return err
	}
	s.collection = data.(*pb.RecordCollection)
	return nil
}

func (s *Server) saveRecordCollection() {
	s.KSclient.Save(KEY, s.collection)
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterDiscogsServiceServer(server, s)
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Mote promotes/demotes this server
func (s *Server) Mote(master bool) error {
	if master {
		return s.readRecordCollection()
	}

	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{&pbg.State{Key: "last_sync_time", TimeValue: s.lastSyncTime.Unix()}}
}

// Init builds out a server
func Init() *Server {
	return &Server{GoServer: &goserver.GoServer{}, lastSyncTime: time.Now()}
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	server := Init()
	server.Register = server
	server.PrepServer()

	server.GoServer.KSclient = *keystoreclient.GetClient(server.GetIP)
	server.RegisterServer("recordcollection", false)
	server.RegisterRepeatingTask(server.runSync, time.Hour)
	server.Log("Starting!")
	server.Serve()
}