//go:generate go-bindata -pkg main -o impl.go  index.js
package main

import (
	"bytes"
	"net"
	"os/exec"

	"github.com/kildevaeld/notto"
	"github.com/kildevaeld/notto/modules"
	"github.com/kildevaeld/notto/modules/archive"
	"github.com/kildevaeld/notto/modules/docker"
	"github.com/kildevaeld/notto/modules/s3"
	homedir "github.com/mitchellh/go-homedir"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {

			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {

	vm := notto.New()

	modules.Define(vm)
	if err := docker.Define(vm, nil); err != nil {
		panic(err)
	}
	if err := s3.Define(vm); err != nil {
		panic(err)
	}

	archive.Define(vm)

	vm.Set("dockermachine", func() string {
		c := exec.Command("docker-machine", "ip")
		out := bytes.NewBuffer(nil)
		c.Stdout = out
		c.Run()
		return string(out.Bytes())
	})

	vm.Set("$HOST_IP", GetLocalIP())
	home, err := homedir.Dir()
	if err == nil {
		vm.Set("$HOME", home)
	}

	_, err = vm.RunScript(string(MustAsset("index.js")), ".")
	if err != nil {
		panic(err)
	}
}
