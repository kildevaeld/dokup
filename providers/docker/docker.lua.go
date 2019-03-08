package docker

import "github.com/aarzilli/golua/lua"

func initLua(l *lua.State) error {

	l.PushGoFunction(func(state *lua.State) int {

		return 0
	})

	return nil
}
