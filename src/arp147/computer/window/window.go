package window

import (
	"arp147/computer/line"
	"engo.io/ecs"
)

type Window struct {
	Active   bool
	Lines    map[int]*line.Line
	Line     int
	Terminal *Terminal

	entity *ecs.Entity
	world  *ecs.World
}

func New(world *ecs.World) *Window {
	w := &Window{
		Active: false,
		Lines:  make(map[int]*line.Line),
		Line:   0,
		world:  world,
	}

	w.Terminal = NewTerminal(world, w)
	return w
}
