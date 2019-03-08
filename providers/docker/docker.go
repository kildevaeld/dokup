package docker

import (
	"context"
	"errors"
	"io"
	"os"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/kildevaeld/dokup"
)

type DockerServiceDescription struct {
	Dockerfile string
	Context    string
	Steps      []string
	Image      string
	Cmd        []string
	Entrypoint []string
	WorkingDir string
	Env        []string
	// Settings
	AttachStderr        bool
	AttachStdout        bool
	AttachStdin         bool
	OpenStdin           bool
	StdinOnce           bool
	AutoRemove          bool
	PublishAllPorts     bool
	Privileged          bool
	TTY                 bool
	Link                []string
	Publish             []string
	RemoveTempContainer bool
	Pull                bool
}

type DockerProvider struct {
	config DockerServiceDescription
	client *docker.Client
	name   string
}

func (d *DockerProvider) ShouldBuild() bool {
	if d.config.Dockerfile != "" || d.config.Steps != nil {
		return true
	}

	return false
}

func getBuildOptions(name string, o DockerServiceDescription, ctx context.Context) (docker.BuildImageOptions, error) {
	var (
		input io.Reader
		//output io.Writer
		err error
	)
	out := docker.BuildImageOptions{}

	if input, err = getDockerContext(o, ctx); err != nil {
		return out, err
	}

	out.InputStream = input
	out.OutputStream = os.Stdout
	out.RmTmpContainer = o.RemoveTempContainer
	out.ForceRmTmpContainer = o.RemoveTempContainer
	out.Pull = o.Pull
	out.Context = ctx

	out.Name = name

	return out, nil
}

func (d *DockerProvider) Build(ctx context.Context) error {

	options, err := getBuildOptions(d.name, d.config, ctx)

	if err != nil {
		return err
	}

	return d.client.BuildImage(options)

}
func (d *DockerProvider) IsRunning() bool {
	i, err := d.client.InspectContainer(d.name)
	if err != nil {
		return false
	} else if i.State.Running {
		return true
	}
	return false
}

func (d *DockerProvider) hasContainer(name string) bool {
	containers, err := d.client.ListContainers(docker.ListContainersOptions{
		All: true,
	})
	if err != nil {

		return false
	}

	for _, i := range containers {
		if name == i.ID {
			return true
		}
		for _, t := range i.Names {
			if t[1:] == name || t == name {
				return true
			}
		}
	}
	return false
}

func (d *DockerProvider) Start() error {

	if d.IsRunning() {
		return errors.New("already running")
	}

	var (
		//container *docker.Container
		options docker.CreateContainerOptions
		err     error
		name    string
	)

	name = d.name

	if options, err = getCreateOptions(d.config, name, name); err != nil {
		return err
	}
	if !d.hasContainer(name) {

		if _, err = d.client.CreateContainer(options); err != nil {
			return err
		}
	} else {
		if err = d.client.StartContainer(name, options.HostConfig); err != nil {
			return err
		}
	}

	return nil
}

func init() {

	dokup.RegisterServiceProvider("docker", func(desc *dokup.ServiceDescription) (dokup.Service, error) {

		var o DockerServiceDescription
		switch t := desc.Config.(type) {
		case *DockerServiceDescription:
			o = *t
		case DockerServiceDescription:
			o = t
		default:
			return nil, errors.New("invalid config")
		}

		client, err := docker.NewClientFromEnv()
		if err != nil {
			return nil, err
		}
		provider := &DockerProvider{
			config: o,
			client: client,
			name:   desc.Name,
		}

		return provider, nil
	})

	if err := dokup.RegisterConfigEntry("docker", &DockerServiceDescription{}); err != nil {
		panic(err)
	}

}
