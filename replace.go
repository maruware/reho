package main

import (
	"fmt"
	"strings"
)

func ReplaceHost(src string, dstHost string) (string, error) {
	s, err := ParseAmbiguousURL(src)
	if err != nil {
		return "", fmt.Errorf("Bad src URL: %s", src)
	}

	d, err := ParseAmbiguousURL(dstHost)
	if err != nil {
		return "", fmt.Errorf("Bad dst host: %s", dstHost)
	}
	if len(d.u.Host) <= 0 {
		return "", fmt.Errorf("Bad dst host: %s", dstHost)
	}

	if len(d.u.Scheme) > 0 {
		s.u.Scheme = d.u.Scheme
	}
	s.u.Host = d.u.Host

	r := s.u.String()
	if s.corrected {
		return strings.TrimPrefix(r, "//"), nil
	}

	return r, nil
}
