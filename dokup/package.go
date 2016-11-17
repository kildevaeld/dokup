package dokup

type PackageDesc struct {
	Name     string                 `json:"name"`
	Services map[string]ServiceDesc `json:"services"`
}

type ServiceDesc struct {
	Name    string   `json:"name"`
	image   string   `json:"image,omitempty"`
	publish []string `json:"publish,omitempty"`
	volume  []string `json:"volume,omitempty`
	link    []string `json:"link,omitempty"`
	Tty     bool     `json:"tty,omitempty"`
	Dependencies
	Production  *ServiceDesc `json:"$production,omitempty"`
	Staging     *ServiceDesc `json:"$staging,omitempty"`
	Development *ServiceDesc `json:"$development,omitempty"`
	Testing     *ServiceDesc `json:"$testing,omitempty"`
	Darwin      *ServiceDesc `json:"$darwin,omitempty"`
	Linux       *ServiceDesc `json:"$linux,omitempty"`
}
