package gohttp

import (
	"net/url"

	"github.com/krishpranav/wpscan/pkg/handler"
)

func IsURL(URL string) bool {
	defer handler.HandleErrorURL()

	uri, err := url.ParseRequestURI(URL)

	if err != nil {
		panic(err)
	}

	switch uri.Scheme {
	case "http":
	case "https":
	default:
		panic("Invalid scheme")

	}

	return true
}
