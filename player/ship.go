package player

import (
	"github.com/lukevers/arp147/terminal"
)

type Ship struct {
	Terminal *terminal.Terminal
}

func NewShip() *Ship {
	return &Ship{
		Terminal: terminal.New(),
	}
}
