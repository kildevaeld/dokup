package dokup

import (
	"net"

	"github.com/Sirupsen/logrus"
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
			} else if ipnet.IP.To16() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func initVm() (*notto.Notto, error) {

	vm := notto.New()

	if err := modules.Define(vm); err != nil {
		return nil, err
	}
	if err := docker.Define(vm, nil); err != nil {
		return nil, err
	}
	if err := s3.Define(vm); err != nil {
		return nil, err
	}

	if err := archive.Define(vm); err != nil {
		return nil, err
	}

	vm.Set("$HOST_IP", GetLocalIP())
	home, err := homedir.Dir()
	if err == nil {
		vm.Set("$HOME", home)
	} else {

		logrus.Warnf("Could determine home directory")
		vm.Set("$HOME", "")
	}
	return vm, nil
}
