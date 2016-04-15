package newgame

import (
	"arp147/computer"
	"arp147/systems/key"
	"arp147/ui/text"
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
	w.AddSystem(&key.KeySystem{})
	w.AddSystem(&text.TextControlSystem{})
	c.CreateEntityUI(w)
	w.AddEntity(c.Entity())
}
