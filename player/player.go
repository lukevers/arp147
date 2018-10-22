package player

import (
	"github.com/lukevers/arp147/clock"
)

type Player struct {
	Clock *clock.Clock
	Ship  *Ship
}

func New() *Player {
	clock := clock.New()
	return &Player{
		Clock: clock,
		Ship:  NewShip(clock),
	}
}
