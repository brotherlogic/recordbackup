package main

import (
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbg "github.com/brotherlogic/goserver/proto"
	pb "github.com/brotherlogic/recordbackup/proto"
	pbrc "github.com/brotherlogic/recordcollection/proto"
)

const (
	// KEY - where records are backed up
	KEY = "/github.com/brotherlogic/recordbackup/config"
)

//Server main server type
type Server struct {
	*goserver.GoServer
	config *pb.Config
	getter getter
}

// Init builds the server
func Init() *Server {
	s := &Server{
		&goserver.GoServer{},
		&pb.Config{},
		&prodGetter{},
	}
	s.getter = &prodGetter{dial: s.DialMaster}
	return s
}

type getter interface {
	getRecords(ctx context.Context) ([]*pbrc.Record, error)
}

type prodGetter struct {
	dial func(server string) (*grpc.ClientConn, error)
}

func (p prodGetter) getRecords(ctx context.Context) ([]*pbrc.Record, error) {
	conn, err := p.dial("recordcollection")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbrc.NewRecordCollectionServiceClient(conn)
	req := &pbrc.GetRecordsRequest{Caller: "recordbackup", Filter: &pbrc.Record{}}
	resp, err := client.GetRecords(ctx, req, grpc.MaxCallRecvMsgSize(1024*1024*1024))
	if err != nil {
		return nil, err
	}

	return resp.GetRecords(), nil
}

func (s *Server) save(ctx context.Context) {
	s.KSclient.Save(ctx, KEY, s.config)
}

func (s *Server) load(ctx context.Context) error {
	config := &pb.Config{}
	data, _, err := s.KSclient.Read(ctx, KEY, config)

	if err != nil {
		return err
	}

	s.config = data.(*pb.Config)
	return nil
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	// Pass
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	if master {
		err := s.load(ctx)
		return err
	}

	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	mapper := make(map[int32]int)
	for _, m := range s.config.Metadata {
		mapper[m.InstanceId]++
	}
	return []*pbg.State{
		&pbg.State{Key: "records", Value: int64(len(mapper))},
		&pbg.State{Key: "stored", Value: int64(len(s.config.Metadata))},
	}
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
	server.GoServer.KSclient = *keystoreclient.GetClient(server.DialMaster)
	server.PrepServer()
	server.Register = server

	err := server.RegisterServer("recordbackup", false)
	if err != nil {
		log.Fatalf("Unable to register: %v", err)
	}

	server.RegisterRepeatingTask(server.procRecords, "proc_records", time.Minute*5)

	server.Serve()
}
