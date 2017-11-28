package player

import (
	"github.com/lukevers/arp147/player/terminal"
)

type Ship struct {
	Terminal *terminal.Terminal
}

func NewShip() *Ship {
	return &Ship{
		Terminal: terminal.New(),
	}
}
