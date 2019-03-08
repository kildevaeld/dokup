package dokup

import (
	"fmt"

	"golang.org/x/net/context"
)

type service struct {
	p Service
	c ServiceDescriptionProvider
}

func (s *service) Description() ServiceDescription {
	return *s.c.Description()
}

func (s *service) Build(ctx context.Context) error {
	if !s.p.ShouldBuild() {
		fmt.Printf("should no build")
		return nil
	}

	s.c.RunHook(PrebuildHook, ctx)

	if err := s.p.Build(ctx); err != nil {
		return err
	}

	s.c.RunHook(BuildHook, ctx)

	return nil
}

/*
func (s *service) RunHook(hook Hook, ctx context.Context) error {
	s.c.RunHook(hook, ctx)
	return nil
}*/

func (s *service) IsRunning() bool {
	return s.p.IsRunning()
}

func (s *service) Start() error {

	return s.p.Start()
}
