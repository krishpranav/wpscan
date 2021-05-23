package gohttp

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	firewallPassing bool
)

type httpOptions struct {
	url                  *URLOptions
	method               string
	tlsCertificateVerify bool
	tor                  bool
	proxy                func(*http.Request) (*url.URL, error)
	data                 io.Reader
	userAgent            string
	redirect             func(req *http.Request, via []*http.Request) error
	contentType          string
	firewall             bool
	sleep                time.Duration
}

type URLOptions struct {
	Simple    string
	Full      string
	Directory string
	URL       *url.URL
}
