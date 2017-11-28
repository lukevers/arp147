package programs

import (
	"engo.io/engo"
	"github.com/lukevers/arp147/input"
	"github.com/lukevers/arp147/terminal/sdk/flag"
	"strings"
)

type Echo struct {
	err error
	out string

	flags *flag.Set
}

func (e *Echo) Name() string {
	return "echo"
}

func (e *Echo) Headless() bool {
	return true
}

func (e *Echo) Error() error {
	return e.err
}

func (e *Echo) Output() string {
	return e.out
}

func (e *Echo) Init() error {
	e.flags = flag.Register([]flag.Flag{
		flag.Flag{
			Long:  "help",
			Short: "h",
		},
	})

	return nil
}

func (e *Echo) Run(args []string) uint {
	// Parse arguments.
	if e.err = e.flags.Parse(args); e.err != nil {
		return 1
	}

	// TODO: use arguments

	e.out = strings.Join(args, " ")
	return 0
}

func (e *Echo) HandleKey(key engo.Key, mods *input.Modifiers) {
	// No need to handle key input since this command doesn't capture/modify
	// the window at all.
}
