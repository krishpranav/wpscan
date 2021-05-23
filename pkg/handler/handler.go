package handler

import (
	"fmt"
	"strings"

	"github.com/krishpranav/wpscan/pkg/printer"
)

func HandlerErrorTorProxy() {
	if recovered := recover(); recovered != nil {
		recoveredS := fmt.Sprintf("%s", recovered)

		if strings.Contains(recoveredS, "proxyconnect tcp: dial tcp 127.0.0.1:9080: connect: connection refused") {
			printer.Fatal("Connection Refused, the tor with the command: \"tor --HTTPTunnelPort 9080\"\n")
		} else {
			printer.Danger(recoveredS).L()
		}
	}
}

func HandlerErrorURL() {
	if recovered := recover(); recovered != nil {
		recoverdS := fmt.Sprintf("%s", recovered)

		printer.Fatal(recoverdS + "\n")
	}
}
