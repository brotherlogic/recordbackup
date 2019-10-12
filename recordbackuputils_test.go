package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/brotherlogic/keystore/client"
	pbrc "github.com/brotherlogic/recordcollection/proto"
)

type testGetter struct {
	fail    bool
	failGet bool
}

func (t *testGetter) getRecords(ctx context.Context, since int64) ([]int32, error) {
	if t.fail {
		return make([]int32, 0), fmt.Errorf("Built to fail")
	}

	return []int32{int32(100)}, nil
}

func (t *testGetter) getRecord(ctx context.Context, id int32) (*pbrc.Record, error) {
	if t.failGet {
		return nil, fmt.Errorf("Built to fail")
	}
	return &pbrc.Record{Metadata: &pbrc.ReleaseMetadata{InstanceId: 100, Cost: 100, DateAdded: 1}}, nil
}

func InitTestServer() *Server {
	s := Init()
	s.SkipLog = true
	s.GoServer.KSclient = *keystoreclient.GetTestClient("./testing")
	s.getter = &testGetter{}
	return s
}

func TestFullMatchFailOnMissingDateAdded(t *testing.T) {
	s := InitTestServer()
	match1 := s.fullMatch(context.Background(), &pbrc.ReleaseMetadata{}, &pbrc.ReleaseMetadata{DateAdded: 1})
	if !match1 {
		t.Errorf("Failed to pass a missing Date")
	}
	match2 := s.fullMatch(context.Background(), &pbrc.ReleaseMetadata{DateAdded: 1}, &pbrc.ReleaseMetadata{})
	if !match2 {
		t.Errorf("Failed to pass a missing Date")
	}

}

func TestFullMatchFailOnDifferentIds(t *testing.T) {
	s := InitTestServer()
	match1 := s.fullMatch(context.Background(), &pbrc.ReleaseMetadata{DateAdded: 1}, &pbrc.ReleaseMetadata{InstanceId: 1, DateAdded: 1})
	if match1 {
		t.Errorf("Matched on diff ids")
	}
	match2 := s.fullMatch(context.Background(), &pbrc.ReleaseMetadata{InstanceId: 1, DateAdded: 1}, &pbrc.ReleaseMetadata{DateAdded: 2})
	if match2 {
		t.Errorf("Matched on diff ids")
	}

}

func TestStandardMatch(t *testing.T) {
	s := InitTestServer()
	err := s.procRecords(context.Background())
	if err != nil {
		t.Errorf("Error processing records: %v", err)
	}

	if len(s.config.Metadata) != 1 {
		t.Errorf("Record was not processed")
	}
}

func TestStandardMatchWithGetFail(t *testing.T) {
	s := InitTestServer()
	s.getter = &testGetter{failGet: true}
	err := s.procRecords(context.Background())
	if err == nil {
		t.Errorf("Error processing records: %v", err)
	}
}

func TestStandardMatchOnFailGet(t *testing.T) {
	s := InitTestServer()
	s.getter = &testGetter{fail: true}
	err := s.procRecords(context.Background())
	if err == nil {
		t.Errorf("Proc did not fail")
	}
}

func TestStandardMatchFindMatch(t *testing.T) {
	s := InitTestServer()
	s.config.Metadata = append(s.config.Metadata, &pbrc.ReleaseMetadata{InstanceId: 100, Cost: 100, DateAdded: 1})
	err := s.procRecords(context.Background())
	if err != nil {
		t.Errorf("Error processing records: %v", err)
	}

	if len(s.config.Metadata) != 1 {
		t.Errorf("Record was not processed")
	}
}

func TestStandardMatchFindMisMatch(t *testing.T) {
	s := InitTestServer()
	s.config.Metadata = append(s.config.Metadata, &pbrc.ReleaseMetadata{InstanceId: 100, Cost: 50, DateAdded: 1})
	err := s.procRecords(context.Background())
	if err != nil {
		t.Errorf("Error processing records: %v", err)
	}

	if len(s.config.Metadata) != 2 {
		t.Errorf("Record was not processed")
	}
}
