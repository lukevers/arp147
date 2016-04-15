package computer

import (
	"arp147/ui/text"
	"engo.io/ecs"
)

type Computer struct {
	world  *ecs.World
	entity *ecs.Entity
	lines  map[int][]*text.Text
	line   int
}

func New(world *ecs.World) *Computer {
	return &Computer{
		world: world,
		lines: make(map[int][]*text.Text),
	}
}
