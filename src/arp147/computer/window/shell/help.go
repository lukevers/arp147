package shell

import (
	"arp147/computer/line"
)

type Help struct {
	Shell
}

func (h *Help) Name() string {
	return "help"
}

func (t *Help) Description() string {
	return "Show a list of commands available with short descriptions"
}

func (t *Help) Help() string {
	return "abc"
}

// This is done at the higher level of the terminal and should be left empty.
// Nothing here will ever run.
func (t *Help) Run(command *line.Command) error {
	return nil
}
