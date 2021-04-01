package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceHost(t *testing.T) {
	testcases := []struct {
		url    string
		dhost  string
		expect string
	}{
		{url: "google.com", dhost: "example.com", expect: "example.com"},
		{url: "http://google.com", dhost: "example.com", expect: "http://example.com"},
		{url: "https://google.com", dhost: "example.com", expect: "https://example.com"},
		{url: "https://my.domain.com/some/path", dhost: "http://localhost:8080", expect: "http://localhost:8080/some/path"},
		{url: "s3://example/key", dhost: "other", expect: "s3://other/key"},
	}

	for _, tt := range testcases {
		t.Run("Replace host of "+tt.url, func(t *testing.T) {
			r, err := ReplaceHost(tt.url, tt.dhost)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expect, r)
		})
	}
}
