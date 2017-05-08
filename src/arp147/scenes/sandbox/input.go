package sandbox

import (
	"arp147/input"
	"arp147/logging"
	"engo.io/ecs"
	"engo.io/engo"
)

func (scene *SandboxScene) SetupInput(world *ecs.World) {
	world.AddSystem(&input.InputSystem{})

	input.RegisterButtons([]input.Key{
		input.Key{
			Name: "action",
			Keys: []engo.Key{engo.Space},
			JustPressed: func() {
				logging.Stdout.Println("just pressed")
			},
			Down: func() {
				logging.Stdout.Println("down")
			},
			JustReleased: func() {
				logging.Stdout.Println("just released")
			},
		},
	})
}
