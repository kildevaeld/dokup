package dokup

import (
	"errors"
	"io"

	"github.com/kildevaeld/notto"
	"github.com/robertkrimen/otto"
)

type Module map[string]interface{}

type Dokup struct {
	vm *notto.Notto
	o  *otto.Object
}

func (self *Dokup) Build() {

}

func (self *Dokup) Start() {

}

func NewDokup(reader io.Reader, ext string) (*Dokup, error) {
	var (
		v   otto.Value
		err error
		vm  notto.Notto
	)

	if vm, err = initVm(); err != nil {
		return nil, err
	}

	switch ext {
	case ".js":
		v, err = formJavascript(vm, reader)
	case ".json":
		v, err = fromJSON(vm, reader)
	default:
		return nil, errors.New("not a known fomat")
	}

	if err != nil {
		return nil, err
	}

}
