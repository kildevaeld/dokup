package dokup

import "context"

type Hook int

const (
	PrestartHook Hook = iota
	PrebuildHook
	PreremoveHook
	StartHook
	BuildHook
	RemoveHook
)

type ServiceDescription struct {
	Name         string
	Type         string
	Config       interface{}
	Dependencies []string
}

type Service interface {
	ShouldBuild() bool
	Build(ctx context.Context) error
	IsRunning() bool
	Start() error
}

type ServiceDescriptionProvider interface {
	Name() string
	Description() *ServiceDescription
	RunHook(hook Hook, ctx context.Context)
}

type ConfigProvider interface {
	LoadFile() error
	Services() []ServiceDescriptionProvider
	Close() error
}

type ServiceFactory func(desc *ServiceDescription) (Service, error)
type ConfigProviderFactory func(path string) (ConfigProvider, error)

var serviceFactories map[string]ServiceFactory
var configProviders map[string]ConfigProviderFactory
var configEntries map[string]ConfigEntryCreator

func RegisterServiceProvider(name string, provider ServiceFactory) {
	serviceFactories[name] = provider
}

func RegisterConfigProvider(ext string, provider ConfigProviderFactory) {
	configProviders[ext] = provider
}

func init() {
	serviceFactories = make(map[string]ServiceFactory)
	configProviders = make(map[string]ConfigProviderFactory)
	configEntries = make(map[string]ConfigEntryCreator)
}
