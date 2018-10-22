package player

import (
	"github.com/lukevers/arp147/clock"
	"github.com/lukevers/arp147/terminal"
)

type Ship struct {
	Clock    *clock.Clock
	Terminal *terminal.Terminal
}

func NewShip(clock *clock.Clock) *Ship {
	return &Ship{
		Clock:    clock,
		Terminal: terminal.New(clock),
	}
}
