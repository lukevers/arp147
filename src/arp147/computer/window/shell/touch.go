package shell

import (
	"arp147/computer/line"
)

type Touch struct {
	Shell
}

func (t *Touch) Name() string {
	return "touch"
}

func (t *Touch) Description() string {
	return "TODO"
}

func (t *Touch) Help() string {
	return "TODO"
}

func (t *Touch) Run(command *line.Command) error {
	// TODO
	return nil
}
