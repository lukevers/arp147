package programs

import (
	"engo.io/engo"
	"github.com/lukevers/arp147/input"
)

type Echo struct{}

func (e Echo) Name() string {
	return "echo"
}

func (e Echo) Run(args []string) int {
	// TODO
	return 0
}

func (e Echo) HandleKey(key engo.Key, mods *input.Modifiers) {
	// No need to handle key input since this command doesn't capture/modify
	// the window at all.
}
