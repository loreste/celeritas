package celeritas

const version = "0.0.1"

type Celeritas struct {
	AppName string
	Debug   bool
	Version string
}

func (c *Celeritas) New(rootPath string) error {
	return nil
}
