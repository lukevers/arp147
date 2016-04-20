package computer

import (
	"arp147/ui/text"
	"engo.io/ecs"
)

type Computer struct {
	world  *ecs.World
	entity *ecs.Entity
	lines  map[int]*line
	line   int
}

type line struct {
	text   []*text.Text
	locked bool
}

func New(world *ecs.World) *Computer {
	return &Computer{
		world: world,
		lines: make(map[int]*line),
	}
}
