package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/danni-popova/analyser/proto"
)

func main() {
	log.Println("starting gRPC server")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialise service
	s := analyserService{}
	grpcServer := grpc.NewServer()

	pb.RegisterAnalyserServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
