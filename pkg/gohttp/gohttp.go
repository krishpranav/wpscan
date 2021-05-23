package gohttp

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/krishpranav/wpscan/internal/database"
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

func SimpleRequest(params ...string) *Response {
	http := NewHTTPClient()
	http.SetURL(params[0])

	if len(params) > 1 {
		http.SetURLDirectory(params[1])
	}

	http.OnTor(database.Memory.GetBool("HTTP Options TOR"))
	http.OnRandomUserAgent(database.Memory.GetBool("HTTP Options Random Agent"))
	http.OnTLSCertificateVerify(database.Memory.GetBool("HTTP Options TLS Certificate Verify"))
	http.FirewallDetection(true)

	response, err := http.Run()

	if err != nil {
		printer.Fatal(err)
	}

	return response
}

func NewHTTPClient() *httpOptions {
	options := &httpOptions{
		method:      "GET",
		userAgent:   "WPrecon - Wordpress Recon (Vulnerability Scanner)",
		data:        nil,
		contentType: "text/html; charset=UTF-8"}

	options.url = &URLOptions{}

	return options
}

func (options *httpOptions) SetURL(url string) *httpOptions {
	if !strings.HasSuffix(url, "/") {
		options.url.Simple = url + "/"
		options.url.Full = url + "/"
	} else {
		options.url.Simple = url
		options.url.Full = url
	}
}
