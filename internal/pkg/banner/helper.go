package banner

import (
	"fmt"
	"runtime"

	color "github.com/logrusorgru/aurora"
)

func green(s string) string {
	if runtime.GOOS == "windows" {
		return s
	}

	return color.Green(s).String()
}

var HelpMain = fmt.Sprintf(`wpscan (Wordpress scanner is a tool for finding wordpress vulnerabilities`)
