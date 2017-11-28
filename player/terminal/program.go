package terminal

import (
	"engo.io/engo"
	"github.com/lukevers/arp147/input"
	"github.com/lukevers/arp147/player/terminal/programs"
)

var (
	registeredPrograms []Program
)

type Program interface {
	Name() string
	Run([]string) int
	HandleKey(engo.Key, *input.Modifiers)
}

func init() {
	registeredPrograms = []Program{
		&programs.Echo{},
	}
}
