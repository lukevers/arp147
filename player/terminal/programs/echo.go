package programs

import (
	"engo.io/engo"
	"github.com/lukevers/arp147/input"
	"strings"
)

type Echo struct {
	err error
	out string
}

func (e *Echo) Error() error {
	return e.err
}

func (e *Echo) Output() string {
	return e.out
}

func (e *Echo) Name() string {
	return "echo"
}

func (e *Echo) Run(args []string) int {
	e.out = strings.Join(args, " ")
	return 0
}

func (e *Echo) HandleKey(key engo.Key, mods *input.Modifiers) {
	// No need to handle key input since this command doesn't capture/modify
	// the window at all.
}
