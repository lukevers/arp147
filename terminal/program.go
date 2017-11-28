package terminal

import (
	"engo.io/engo"
	"github.com/lukevers/arp147/input"
	"github.com/lukevers/arp147/terminal/programs"
)

var (
	registeredPrograms []Program
)

// Program defines a shell program that is either a headless program or a UI
// based program.
type Program interface {
	// An error message returned by the program if there are any issues. All
	// programs SHOULD implement this to give the user a descriptive error
	// message.
	Error() error

	// Handle a key press given to the program. Headless programs MUST NOT
	// implement this, and UI based programs MUST implement this.
	HandleKey(engo.Key, *input.Modifiers)

	// Defines if a program is headless or UI based.
	Headless() bool

	// A function that's called before Run is called. All functions SHOULD
	// implement this. If an error is returned, Run will not be called.
	Init() error

	// The callable name of the program.
	Name() string

	// An output message returned by the program. Headless programs SHOULD
	// implement this to give the user a descriptive output, although it is not
	// required, and most likely depends on the context of the program. UI
	// based programs SHOULD NOT implement this.
	Output() string

	// Run the program with the arguments given. All programs MUST implement
	// this. An error code (0 being OK, > 0 being ERROR) MUST be returned. For
	// headless programs, this will determine if Error() or Output() is called
	// on the program.
	Run([]string) uint
}

func init() {
	registeredPrograms = []Program{
		&programs.Echo{},
	}
}
