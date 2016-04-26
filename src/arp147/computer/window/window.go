package window

import (
	"arp147/computer/line"
	"engo.io/ecs"
)

type Window struct {
	Active bool
	Lines  map[int]*line.Line
	Line   int

	entity *ecs.Entity
	world  *ecs.World
}

func New(world *ecs.World) *Window {
	return &Window{
		Active: false,
		Lines:  make(map[int]*line.Line),
		Line:   0,
		world:  world,
	}
}
