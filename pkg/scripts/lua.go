package scripts

var (
	PathScript = "tools/scripts"
)

type structscript struct {
	Title       string
	Author      string
	License     string
	Description string
	References  []string
	RiskLevel   string
}

type script struct {
	lstate *lua.LState
}
