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

func NewScript() *script {
	LuaNewState := lua.NewState()

	LuaNewState.PreloadModule("url", luaUrl.Loader)
	LuaNewState.PreloadModule("printer", luaPrinter.Loader)
	LuaNewState.PreloadModule("net", luaNet.Loader)

	luaLibs.Preload(LuaNewState)

	LuaNewState.SetGlobal("tor_url", lua.LString("http://127.0.0.1:9080"))

	return &script{lstate: LuaNewState}
}
