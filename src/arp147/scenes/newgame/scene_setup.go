package newgame

import (
	"arp147/computer"
	"arp147/systems/key"
	"engo.io/ecs"
	"engo.io/engo"
	"image/color"
)

func (s *Scene) Setup(w *ecs.World) {
	engo.SetBackground(color.Black)

	w.AddSystem(&engo.MouseSystem{})
	w.AddSystem(&engo.RenderSystem{})
	w.AddSystem(&key.KeySystem{})

	// -- Controls

	controls := ecs.NewEntity("KeySystem")
	kcontrol := &key.KeyComponent{}

	// -- Computer

	c := computer.New(w)
	w.AddEntity(c.Entity())
	kcontrol.On(engo.C, func(key engo.Key) {
		if !c.Active {
			c.StartSession()
		}
	})

	// -- Controls

	controls.AddComponent(kcontrol)
	w.AddEntity(controls)
}
