package dokup

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/Sirupsen/logrus"
	multierror "github.com/hashicorp/go-multierror"
)

type Dokup struct {
	services map[string]*service
}

func (d *Dokup) LoadDescriptionFile(path string) error {
	ext := filepath.Ext(path)

	var (
		conf ConfigProvider
		desc *ServiceDescription
		pro  Service
		err  error
	)

	factory, ok := configProviders[ext[1:]]

	if !ok {
		return errors.New("no config provider for: " + path)
	}

	if conf, err = factory(path); err != nil {
		return err
	}

	if err = conf.LoadFile(); err != nil {
		return err
	}

	for _, s := range conf.Services() {

		if desc = s.Description(); desc == nil {
			return errors.New("no description")
		}

		profac := serviceFactories[desc.Type]
		if profac == nil {
			return errors.New("service not found")
		}

		if pro, err = profac(desc); err != nil {
			return err
		}

		d.services[desc.Name] = &service{
			p: pro,
			c: s,
		}

	}

	return nil

}

func (d *Dokup) resolveDependencies(name string, cache map[string]*service, parent string) ([]*service, error) {
	serv := d.services[name]
	deps := serv.Description().Dependencies
	if deps == nil {
		return nil, nil
	}

	if cache == nil {
		cache = make(map[string]*service)
	}

	var out []*service
	for _, dep := range deps {
		dserv := d.services[dep]
		if dserv == nil {
			return nil, fmt.Errorf("could not find dependency: %s for %s", dep, name)
		}
		if dep == parent {
			return nil, fmt.Errorf("cirkel")
		} else if _, ok := cache[dep]; ok {
			continue
		}
		out = append(out, dserv)
		if dservs, err := d.resolveDependencies(dep, cache, name); dservs != nil {
			out = append(out, dservs...)
		} else if err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (d *Dokup) Build(name string) error {
	ctx := context.Background()
	var lock sync.Mutex
	var result error
	var wg sync.WaitGroup

	addError := func(e error) {
		lock.Lock()
		defer lock.Unlock()
		result = multierror.Append(result, e)
	}

	build := func(s *service) {
		defer wg.Done()
		if err := s.Build(ctx); err != nil {
			addError(err)
		}
	}

	serv, ok := d.services[name]
	if !ok {
		return errors.New("service not found")
	}

	if deps, err := d.resolveDependencies(name, nil, ""); deps != nil {

		for _, dep := range deps {
			wg.Add(1)
			go build(dep)
		}

	} else if err != nil {
		return err
	}
	wg.Add(1)
	go build(serv)
	wg.Wait()

	return result
}

func (d *Dokup) Start(name string) error {
	logrus.Infof("starting: %s", name)
	if deps, err := d.resolveDependencies(name, nil, ""); deps != nil {

		for _, dep := range deps {
			logrus.Infof("starting dependency: %s", dep.Description().Name)
			if err := dep.Start(); err != nil {
				return err
			}

		}

	} else if err != nil {
		return err
	}
	service, ok := d.services[name]
	if !ok {
		return errors.New("service not found")
	}

	return service.Start()

}

func New() *Dokup {
	return &Dokup{
		services: make(map[string]*service),
	}
}
