package gohttp

import (
	"math/rand"
	"time"

	"github.com/krishpranav/wpscan/pkg/wordlist"
)

func randomuseragent() string {
	timeUnix := time.Now().Unix()

	rand.Seed(timeUnix)
	randomValue := rand.Intn(len(wordlist.UserAgents))

	return wordlist.UserAgents[randomValue]
}
