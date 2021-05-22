package main

import (
	"runtime"

	"github.com/krishpranav/wpscan/cli"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cli.Execute()
}
