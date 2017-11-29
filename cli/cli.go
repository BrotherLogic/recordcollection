package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
	pbgd "github.com/brotherlogic/godiscogs"
	"github.com/brotherlogic/goserver/utils"
	pbrc "github.com/brotherlogic/recordcollection/proto"
)

func findServer(name string) (string, int32) {
	conn, _ := grpc.Dial(utils.Discover, grpc.WithInsecure())
	defer conn.Close()

	registry := pbdi.NewDiscoveryServiceClient(conn)
	rs, err := registry.Discover(context.Background(), &pbdi.RegistryEntry{Name: name})

	if err == nil {
		return rs.GetIp(), rs.GetPort()
	}

	return "", -1
}

func main() {
	switch os.Args[1] {
	case "get":
		host, port := findServer("recordcollection")

		conn, _ := grpc.Dial(host+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
		defer conn.Close()

		registry := pbrc.NewDiscogsServiceClient(conn)
		i, _ := strconv.Atoi(os.Args[2])
		rec, err := registry.GetRecords(context.Background(), &pbrc.GetRecordsRequest{Filter: &pbrc.Record{Release: &pbgd.Release{Id: int32(i)}}})

		if err == nil {

			if len(rec.GetRecords()) == 0 {
				log.Printf("No records found!")
			}

			for _, r := range rec.GetRecords() {
				fmt.Printf("Release: %v", r.GetRelease())
				fmt.Printf("Metadata: %v", r.GetMetdata())
			}
		} else {
			log.Printf("Error: %v", err)
		}
	}
}