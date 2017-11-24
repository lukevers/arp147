package player

import (
	"github.com/lukevers/arp147/clock"
)

type Player struct {
	Clock *clock.Clock
	Ship  *Ship
}

func New() *Player {
	return &Player{
		Ship:  NewShip(),
		Clock: clock.New(),
	}
}
