//go:generate go-bindata -pkg main -o impl.go  index.js
package main

import (
	"github.com/kildevaeld/notto"
	"github.com/kildevaeld/notto/modules"
	"github.com/kildevaeld/notto/modules/docker"
)

func main() {

	vm := notto.New()

	modules.Define(vm)
	docker.Define(vm, nil)

	_, err := vm.RunScript(string(MustAsset("index.js")), ".")
	if err != nil {
		panic(err)
	}
}
