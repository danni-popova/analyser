package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/danni-popova/analyser/proto"
)

func Test_AnalyserService_Analyse(t *testing.T) {
	r := require.New(t)
	s := analyserService{}
	ctx := context.Background()

	tests := []struct {
		name      string
		userAgent string
		browser   string
		error     string
		errCode   codes.Code
	}{
		// User agent string examples taken from:
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent
		{
			name:      "User agent string is empty",
			userAgent: "",
			error:     ErrorEmptyUserAgentString,
			errCode:   codes.InvalidArgument,
		},
		{
			name:      "User agent browser is Firefox",
			userAgent: "Mozilla/5.0 (Android 4.3; Mobile; rv:54.0) Gecko/54.0 Firefox/54.0",
			browser:   "Firefox",
		},
		{
			name:      "User agent browser is Safari",
			userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari/604.1",
			browser:   "Safari",
		},
		{
			name:      "User agent browser is Chrome",
			userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
			browser:   "Chrome",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := s.AnalyseUserAgent(ctx, &pb.AnalyseUserAgentRequest{UserAgent: test.userAgent})

			switch test.error {
			case "":
				r.Equal(test.browser, resp.Browser)
			default:
				r.Error(err)
				r.Equal(codes.InvalidArgument, status.Code(err))
			}
		})
	}
}
