package printer

import (
	"os"

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
