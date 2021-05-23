package gohttp

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	TorURL = "http://127.0.0.1:9080"
)

func Tor() (func(*http.Request) (*url.URL, error), error) {
	tor, err := url.Parse(TorURL)

	if err != nil {
		return nil, fmt.Errorf("proxy URL is an invalid (%w)", err)
	}

	return http.ProxyURL(tor), nil
}
