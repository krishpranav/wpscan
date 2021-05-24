package net

import (
	"fmt"
	"net"

	"github.com/krishpranav/wpscan/pkg/printer"
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)

	L.Push(mod)
	return 1
}

func lookupip(L *lua.LState) int {
	ips, err := net.LookupIP(L.ToString(1))

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	ip := fmt.Sprintf("%s", ips[0])

	L.Push(lua.LString(ip))

	return 1
}
