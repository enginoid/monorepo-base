package main

//go:generate protoc -I ./proto --go_out=plugins=grpc:./proto ./proto/ping.proto

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/enginoid/monorepo-base/services/ping/proto"
)

const (
	address = "localhost:50051"
)

type server struct {
	pb.UnimplementedPingServer
}

// SayHello implements helloworld.PingServer
func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	log.Printf("received ping: %#v", in.GetMessage())
	return &pb.PingReply{Message: in.GetMessage()}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingServer(s, &server{})

	log.Printf("listening on %s", address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
