package main

import (
	"log"
	"time"

	pb "github.com/brotherlogic/recordcollection/proto"
	"github.com/golang/protobuf/proto"
)

func (s *Server) syncCollection() {
	records := s.retr.GetCollection()

	for _, record := range records {
		found := false
		for _, r := range s.collection.GetRecords() {
			if r.GetRelease().Id == record.Id {
				found = true
				log.Printf("MERGING %v and %v", r.Release, &record)
				proto.Merge(r.Release, &record)
			}
		}
		if !found {
			s.collection.Records = append(s.collection.Records, &pb.Record{Release: &record})
		}
	}
	s.lastSyncTime = time.Now()
}

func (s *Server) syncWantlist() {
	wants := s.retr.GetWantlist()

	for _, want := range wants {
		found := false
		for _, w := range s.collection.GetWants() {
			if w.Id == want.Id {
				found = true
				proto.Merge(w, &want)
			}
		}
		if !found {
			s.collection.Wants = append(s.collection.Wants, &want)
		}
	}
}

func (s *Server) runSync() {
	//SyncWithDiscogs Syncs everything with discogs
	s.syncCollection()
	s.syncWantlist()
	s.saveRecordCollection()
}
