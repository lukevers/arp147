package terminal

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/lukevers/arp147/clock"
	"github.com/lukevers/arp147/input"
	"github.com/lukevers/arp147/player"
	"log"
)

// Scene defines a scene for the main menu
type Scene struct {
	Player *player.Player
}

// Type defines the scene
func (s *Scene) Type() string {
	return "TerminalScene"
}

// Preload
func (s *Scene) Preload() {
	if err := engo.Files.Load(
		"textures/terminal/background.png",
		"textures/terminal/corner_bottom_left.png",
		"textures/terminal/corner_bottom_right.png",
		"textures/terminal/corner_top_left.png",
		"textures/terminal/corner_top_right.png",
		"textures/terminal/horizontal.png",
		"textures/terminal/vertical.png",
	); err != nil {
		panic(err)
	}
}

// Setup
func (s *Scene) Setup(world *ecs.World) {
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&clock.ClockSystem{Clock: s.Player.Clock})
	world.AddSystem(&input.InputSystem{})

	s.Player.Ship.Terminal.AddToWorld(world)

	input.RegisterKeys([]input.Key{
		input.Key{
			Name: "[a-z][0-9]",
			Keys: []engo.Key{
				engo.A,
				engo.B,
				engo.C,
				engo.D,
				engo.E,
				engo.F,
				engo.G,
				engo.H,
				engo.I,
				engo.J,
				engo.K,
				engo.L,
				engo.M,
				engo.N,
				engo.O,
				engo.P,
				engo.Q,
				engo.R,
				engo.S,
				engo.T,
				engo.U,
				engo.V,
				engo.W,
				engo.X,
				engo.Y,
				engo.Z,
				engo.Zero,
				engo.One,
				engo.Two,
				engo.Three,
				engo.Four,
				engo.Five,
				engo.Six,
				engo.Seven,
				engo.Eight,
				engo.Nine,
			},
			OnPress: func(key engo.Key, mods *input.Modifiers) {
				log.Println("==========")
				log.Println("key:\t", key)
				log.Println("ctl:\t", mods.Control)
				log.Println("alt:\t", mods.Alt)
				log.Println("sft:\t", mods.Shift)
				log.Println("sup:\t", mods.Super)
			},
		},
	})
}
