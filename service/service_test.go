package main

import (
	"testing"
)

func Test_AnalyserService_Analyse(t *testing.T) {
	tests := []struct {
		name      string
		userAgent string
	}{
		{
			name:      "",
			userAgent: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
