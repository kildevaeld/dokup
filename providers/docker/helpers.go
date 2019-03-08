package docker

import (
	"archive/tar"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	docker "github.com/fsouza/go-dockerclient"

	"golang.org/x/net/context"
)

func writeFile(tr *tar.Writer, path string, t time.Time) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	tr.WriteHeader(&tar.Header{
		Name:       info.Name(),
		Size:       info.Size(),
		ModTime:    info.ModTime(),
		AccessTime: t,
		ChangeTime: t,
		Mode:       int64(info.Mode()),
	})

	var b []byte

	if b, err = ioutil.ReadFile(path); err != nil {
		return err
	}

	_, err = tr.Write(b)

	return err
}

func getDockerContext(o DockerServiceDescription, ctx context.Context) (io.Reader, error) {
	var err error
	path := o.Context
	if path == "" {
		path = "."
	}
	if !filepath.IsAbs(path) {
		if path, err = filepath.Abs(path); err != nil {
			return nil, err
		}
	}

	t := time.Now()
	inputbuf := bytes.NewBuffer(nil)
	tr := tar.NewWriter(inputbuf)
	defer tr.Close()

	dockerfileFound := false

	if o.Dockerfile != "" {
		p := o.Dockerfile
		if !filepath.IsAbs(p) {
			if p, err = filepath.Abs(p); err != nil {
				return nil, err
			}
		}
		if o.Context == "" {
			path = filepath.Dir(p)
		} else {
			if err = writeFile(tr, p, t); err != nil {
				return nil, err
			}
			dockerfileFound = true
		}

	} else if o.Steps != nil {

		hasFrom := false
		for _, step := range o.Steps {
			l := strings.ToLower(step)
			if strings.HasPrefix(l, "from ") {
				hasFrom = true
			}
		}
		dockerfile := ""
		if !hasFrom {
			dockerfile = strings.Join(append([]string{"FROM " + o.Image}, o.Steps...), "\n")
		} else {
			dockerfile = strings.Join(o.Steps, "\n")
		}

		tr.WriteHeader(&tar.Header{
			Name:       "Dockerfile",
			Size:       int64(len(dockerfile)),
			ModTime:    t,
			AccessTime: t,
			ChangeTime: t,
			Mode:       int64(0777),
		})

		tr.Write([]byte(dockerfile))

		dockerfileFound = true

	}

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	err = filepath.Walk(path, func(p string, file os.FileInfo, err error) error {
		if err != nil || p == path {
			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		fp := p

		p = strings.Replace(p, path+"/", "", 1)

		if file.Name() == "Dockerfile" {
			if dockerfileFound {
				return errors.New("Cannot contain more than one dockerfile")
			}
			dockerfileFound = true
		}

		if file.IsDir() {
			return nil
		}

		return writeFile(tr, fp, t)

	})

	if !dockerfileFound {
		return nil, errors.New("No dockerfiles in " + path)
	}

	return inputbuf, err
}

func getCreateOptions(o DockerServiceDescription, name, image string) (docker.CreateContainerOptions, error) {
	out := docker.CreateContainerOptions{}
	out.Name = name

	cfg := docker.Config{}
	cfg.Image = image

	var err error

	cfg.Cmd = o.Cmd
	cfg.Entrypoint = o.Entrypoint
	cfg.WorkingDir = o.WorkingDir
	cfg.Env = o.Env
	cfg.AttachStderr = o.AttachStderr
	cfg.AttachStdin = o.AttachStdin
	cfg.AttachStdout = o.AttachStdout
	cfg.Tty = o.TTY
	cfg.OpenStdin = o.OpenStdin
	cfg.StdinOnce = o.StdinOnce

	out.Config = &cfg

	hcfg := &docker.HostConfig{}
	hcfg.AutoRemove = o.AutoRemove
	hcfg.PublishAllPorts = o.PublishAllPorts
	hcfg.Privileged = o.Privileged

	out.HostConfig = hcfg

	return out, err
	/*if cfg.Env, err = getEnv(o); err != nil {
		return out, err
	}

	cfg.AttachStderr = boolOr(o, "attachStderr", false)
	cfg.AttachStdin = boolOr(o, "attachStdin", false)
	cfg.AttachStdout = boolOr(o, "attachStdout", false)
	cfg.Tty = boolOr(o, "tty", false)
	cfg.OpenStdin = boolOr(o, "openStdin", false)
	cfg.StdinOnce = boolOr(o, "stdinOnce", false)

	out.Config = &cfg

	hcfg := &dockerclient.HostConfig{}

	if boolOr(o, "restart", false) {
		hcfg.RestartPolicy = dockerclient.RestartUnlessStopped()
	}

	hcfg.AutoRemove = boolOr(o, "autoRemove", false)
	hcfg.PublishAllPorts = boolOr(o, "publishAllPorts", false)
	hcfg.Privileged = boolOr(o, "previleged", false)
	if hcfg.Binds, err = getStringSlice(o, "binds"); err != nil {
		return out, err
	}

	if hcfg.Links, err = getStringSlice(o, "links"); err != nil {
		return out, err
	}

	if hcfg.Binds, err = getStringSlice(o, "volumes"); err != nil {
		return out, err
	}

	if v, e := getStringSlice(o, "publish"); e == nil {

		p := make(map[dockerclient.Port][]dockerclient.PortBinding)
		for _, k := range v {
			s := strings.Split(k, ":")
			if len(s) != 2 {
				return out, errors.New("")
			}
			p[dockerclient.Port(s[0]+"/tcp")] = []dockerclient.PortBinding{dockerclient.PortBinding{
				HostPort: s[1],
				HostIP:   "0.0.0.0",
			}}

		}
		//hcfg.PublishAllPorts = true
		hcfg.PortBindings = p

	}

	out.HostConfig = hcfg
	out.Config.ExposedPorts = map[dockerclient.Port]struct{}{
		"3306/tcp": {}}

	return out, nil*/

}
