package newgame

import (
	"arp147/computer"
	"arp147/ui/background"
	"engo.io/ecs"
	"engo.io/engo"
	"image/color"
)

func (s *Scene) Setup(w *ecs.World) {
	engo.SetBackground(color.Black)

	w.AddSystem(&engo.MouseSystem{})
	w.AddSystem(&engo.RenderSystem{})

	// -- Background

	background.TileWorld(w, background.Background{
		Scale:   engo.Point{1, 1},
		Texture: engo.Files.Image("space.png"),
	})

	// -- Computer

	c := computer.New()
	w.AddSystem(&computer.ComputerSystem{})
	w.AddEntity(c.Entity())
}
