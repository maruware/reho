package main

import "net/url"

type ParseResult struct {
	u         *url.URL
	corrected bool
}

func ParseAmbiguousURL(rawurl string) (ParseResult, error) {
	var result ParseResult = ParseResult{corrected: false}
	u, err := url.Parse(rawurl)
	if err != nil {
		return ParseResult{}, err
	}
	if u.Host == "" {
		u, err = url.Parse("//" + rawurl)
		if err != nil {
			return ParseResult{}, err
		}
		result.corrected = true
	}
	result.u = u

	return result, nil
}
