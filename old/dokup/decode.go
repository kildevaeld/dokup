package dokup

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/kildevaeld/notto"
	"github.com/robertkrimen/otto"
)

func fromJSON(vm *notto.Notto, reader io.Reader) (otto.Value, error) {
	dec := json.NewDecoder(reader)

	var out Module
	if err := dec.Decode(&out); err != nil {
		return otto.UndefinedValue(), err
	}

	return vm.ToValue(out)
}

func formJavascript(vm *notto.Notto, reader io.Reader) (otto.Value, error) {

	b, e := ioutil.ReadAll(reader)
	if e != nil {
		return otto.UndefinedValue(), e
	}

	return vm.RunScript(string(b), ".")

}
