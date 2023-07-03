package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/danni-popova/analyser/proto"
)

type analyserService struct {
	pb.AnalyserServer
}

// AnalyseUserAgent implements an endpoint that accepts a user agent string,
// analyses it and returns information about it.
func (s *analyserService) AnalyseUserAgent(ctx context.Context, request *pb.AnalyseUserAgentRequest) (*pb.AnalyseUserAgentResponse, error) {
	return &pb.AnalyseUserAgentResponse{}, status.Error(codes.Unimplemented, "Unimplemented")
}
