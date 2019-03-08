//go:generate go-bindata -pkg lua  -o lua.prelude.go dokup.lua

package lua

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aarzilli/golua/lua"
	"github.com/kildevaeld/dict"
	"github.com/kildevaeld/dokup"
	"github.com/mitchellh/mapstructure"
	"github.com/stevedonovan/luar"
)

type LuaDescriptionProvider struct {
	d *dokup.ServiceDescription
	h map[dokup.Hook]*luar.LuaObject
}

func (l *LuaDescriptionProvider) Name() string {
	return ""
}
func (l *LuaDescriptionProvider) Description() *dokup.ServiceDescription {
	return l.d
}

func (l *LuaDescriptionProvider) RunHook(hook dokup.Hook, ctx context.Context) {
	if h, ok := l.h[hook]; ok {
		if err := h.Call(nil); err != nil {
			fmt.Printf("error %v", err)
		}
	}
}

func (l *LuaDescriptionProvider) Close() {
	for _, h := range l.h {
		h.Close()
	}
}

type LuaConfigProvider struct {
	path string
	l    *lua.State
	d    []dokup.ServiceDescriptionProvider
}

func (l *LuaConfigProvider) Services() []dokup.ServiceDescriptionProvider {
	return l.d
}

func strToHook(str string) dokup.Hook {
	switch str {
	case "prestart":
		return dokup.PrestartHook
	case "prebuild":
		return dokup.PrebuildHook
	case "preremove":
		return dokup.PreremoveHook
	}
	return -1
}

var reg = regexp.MustCompile("name")

func (l *LuaConfigProvider) registerService(config luar.Map) {

	m := dict.Map(config)
	var options dokup.ServiceDescription
	if err := mapstructure.Decode(config, &options); err != nil {
		panic(err)
	}
	options.Name = m.GetString("name")

	for key, val := range config {
		if reg.Match([]byte(key)) {
			continue
		}
		if dokup.HasConfigEntry(key) {
			o, e := dokup.CreateConfigEntry(key, val.(map[string]interface{}))
			if e != nil {
				fmt.Printf("error %s\n", e)
			}
			options.Type = key
			options.Config = o
		}

	}

	p := &LuaDescriptionProvider{
		d: &options,
		h: make(map[dokup.Hook]*luar.LuaObject),
	}

	l.d = append(l.d, p)

	for _, h := range []string{"prestart", "prebuild", "preremove"} {
		if m, ok := m.Get(h).(*luar.LuaObject); ok {
			p.h[strToHook(h)] = m
		}
	}

}

func (l *LuaConfigProvider) LoadFile() error {

	luar.Register(l.l, "", luar.Map{
		"_registerService": l.registerService,
	})

	return l.l.DoFile(l.path)
}

func (l *LuaConfigProvider) Close() error {
	for _, ll := range l.d {
		ll.(*LuaDescriptionProvider).Close()
	}
	l.d = nil
	return nil
}

func init() {
	dokup.RegisterConfigProvider("lua", func(path string) (dokup.ConfigProvider, error) {

		l := luar.Init()
		l.OpenLibs()
		if err := l.DoString(string(MustAsset("dokup.lua"))); err != nil {
			return nil, err
		}

		return &LuaConfigProvider{
			path: path,
			l:    l,
		}, nil
	})
}
