package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceHost(t *testing.T) {
	testcases := []struct {
		src      string
		dhost    string
		location bool
		expect   string
	}{
		{src: "google.com", dhost: "example.com", expect: "example.com"},
		{src: "http://google.com", dhost: "example.com", expect: "http://example.com"},
		{src: "https://google.com", dhost: "example.com", expect: "https://example.com"},
		{src: "https://my.domain.com/some/path", dhost: "http://localhost:8080", expect: "http://localhost:8080/some/path"},
		{src: "s3://example/key", dhost: "other", expect: "s3://other/key"},
		{src: "https://g.co/cast/power", dhost: "example.com", location: true, expect: "https://example.com/chromecast/answer/10087405"},
	}

	for _, tt := range testcases {
		t.Run("Replace host of "+tt.src, func(t *testing.T) {
			r, err := ReplaceHost(tt.src, tt.dhost, tt.location)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expect, r)
		})
	}
}
