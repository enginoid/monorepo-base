package main

//go:generate protoc -I ./proto --go_out=plugins=grpc:./proto ./proto/ping.proto

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/enginoid/monorepo-base/services/ping/proto"
	"github.com/enginoid/monorepo-base/services/ping/server"
)

const (
	address = "localhost:50051"
)

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingServer(s, server.NewServer())

	log.Printf("listening on %s", address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
