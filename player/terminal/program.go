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
	Error() error
	HandleKey(engo.Key, *input.Modifiers)
	Name() string
	Output() string
	Run([]string) int
}

func init() {
	registeredPrograms = []Program{
		&programs.Echo{},
	}
}
