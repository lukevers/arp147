package newgame

import (
	"arp147/computer"
	"engo.io/ecs"
	"engo.io/engo"
	"image/color"
)

func (s *Scene) Setup(w *ecs.World) {
	engo.SetBackground(color.Black)

	w.AddSystem(&engo.MouseSystem{})
	w.AddSystem(&engo.RenderSystem{})

	// -- Computer

	c := computer.New(w)
	w.AddEntity(c.Entity())
}
