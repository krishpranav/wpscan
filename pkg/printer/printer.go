package printer

import (
	"os"
	"runtime"

	color "github.com/logrusorgru/aurora"
)

var (
	stdin    = *os.Stdin
	stdout   = *os.Stdout
	stderr   = *os.Stderr
	line     = &ln{}
	zfill    = &z{}
	Required = color.Red("(Required)").Bold().String()
)

var (
	prefixDanger  = color.Red("[✗]").String()
	prefixFatal   = color.Red("[!]").String()
	prefixDone    = color.Green("[✔]").String()
	prefixWarning = color.Yellow("[!]").String()
	prefixScan    = color.Yellow("[?]").String()
	prefixInfo    = color.Magenta("[i]").String()

	prefixListDanger  = color.Red("    —").String()
	prefixListDone    = color.Green("    —").String()
	prefixListDefault = color.White("    —").String()
	prefixListWarning = color.Yellow("    —").String()

	prefixTopLine = color.Yellow("[✲]").String()
)

var seekCurrent = 1

func init() {
	if runtime.GOOS == "windows" {
		prefixDanger = "[✗]"
		prefixFatal = "[!]"
		prefixDone = "[✔]"
		prefixWarning = "[!]"
		prefixScan = "[?]"
		prefixInfo = "[i]"

		prefixListDanger = "    —"
		prefixListDone = "    —"
		prefixListDefault = "    —"
		prefixListWarning = "    —"

		prefixTopLine = "[✲]"
	}
}
