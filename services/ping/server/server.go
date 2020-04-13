package server

import (
	"context"
	"log"

	pb "github.com/enginoid/monorepo-base/services/ping/proto"
)

type server struct {
	pb.UnimplementedPingServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	log.Printf("received ping: %#v", in.GetMessage())
	return &pb.PingReply{Message: in.GetMessage()}, nil
}
