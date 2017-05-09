package sandbox

import (
	"arp147/clock"
	"arp147/input"
	"arp147/logging"
	"engo.io/engo"
)

func (scene *SandboxScene) SetupInput() {
	input.RegisterButtons([]input.Key{
		input.Key{
			Name: "action",
			Keys: []engo.Key{engo.Space},
			JustPressed: func() {
				logging.Stdout.Println("just pressed")
			},
			Down: func() {
				logging.Stdout.Println("time:", clock.String())
			},
			JustReleased: func() {
				logging.Stdout.Println("just released")
			},
		},
	})
}
