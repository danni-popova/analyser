package main

import (
	"context"
	"errors"

	"github.com/mileusna/useragent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/danni-popova/analyser/proto"
)

type analyserService struct {
	pb.AnalyserServer
}

const (
	ErrorEmptyUserAgentString = "user agent string is empty"
)

// AnalyseUserAgent implements an endpoint that accepts a user agent string,
// analyses it and returns information about it.
func (s *analyserService) AnalyseUserAgent(ctx context.Context, request *pb.AnalyseUserAgentRequest) (*pb.AnalyseUserAgentResponse, error) {
	ok, err := s.validateRequest(request.UserAgent)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	parsed := useragent.Parse(request.UserAgent)

	return &pb.AnalyseUserAgentResponse{
		Browser: parsed.Name,
	}, nil
}

func (s *analyserService) validateRequest(userAgent string) (bool, error) {
	if userAgent == "" {
		return false, errors.New(ErrorEmptyUserAgentString)
	}
	return true, nil
}
