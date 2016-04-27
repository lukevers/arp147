package shell

import (
	"arp147/computer/line"
)

type Shell interface {
	// A name of the command
	Name() string

	// Meant to be a short description
	Description() string

	// A full-text description intended to help the user how to use the command
	Help() string

	// A func that runs when the command is parsed
	Run(command *line.Command) error
}
