package gohttp

import (
	"net/http"
	"sync"
)

type welapplicationfirewall struct {
	http       *http.Response
	raw        string
	exists     bool
	output     string
	solve      string
	firewall   string
	confidence int
}

var wg sync.WaitGroup

func NewFirewallDetectionPassive(response *Response) *webapplicationfirewall {
	return &webapplicationfirewall{http: response.Response, raw: response.Raw}
}
