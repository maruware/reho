package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAmbiguousURL(t *testing.T) {
	testcases := []struct {
		url     string
		ehost   string
		escheme string
	}{
		{url: "google.com", ehost: "google.com", escheme: ""},
		{url: "http://google.com", ehost: "google.com", escheme: "http"},
		{url: "s3://example/key", ehost: "example", escheme: "s3"},
	}

	for _, tt := range testcases {
		t.Run("Parse "+tt.url, func(t *testing.T) {
			r, err := ParseAmbiguousURL(tt.url)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.ehost, r.u.Host)
			assert.Equal(t, tt.escheme, r.u.Scheme)
		})
	}
}
