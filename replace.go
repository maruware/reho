package main

import (
	"fmt"
	"net/http"
	"strings"
)

func getLocation(src string) (string, error) {
	client := &http.Client{
		// no redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := client.Get(src)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusFound, http.StatusMovedPermanently:
		loc := res.Header.Get("Location")

		if len(loc) >= 0 {
			return loc, nil
		}
		return src, nil
	default:
		return src, nil
	}
}

func ReplaceHost(src string, dstHost string, location bool) (string, error) {
	if location {
		l, err := getLocation(src)
		if err != nil {
			return "", err
		}
		src = l
	}

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
