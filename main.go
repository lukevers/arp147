package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/lukevers/arp147/input"
	"github.com/lukevers/arp147/navigator"
	"github.com/lukevers/arp147/terminal"
	"github.com/lukevers/arp147/ui"
	"github.com/lukevers/arp147/viewers/opposition"
	"github.com/lukevers/arp147/viewers/user"
)

type defaultScene struct{}

// Type uniquely defines the game type.
func (*defaultScene) Type() string { return "defaultScene" }

// Preload is called before loading any assets from the disk, to allow you to
// register / queue them.
func (*defaultScene) Preload() {
	engo.Files.Load(
		"fonts/Undefined.ttf",

		"textures/bkg_t1.jpg",
		"textures/bkg_t2.jpg",
		"textures/bkg_t2_grid.png",

		"textures/user_01.png",

		"textures/enemy_01.png",
		"textures/enemy_02.png",
		"textures/enemy_03.png",
		"textures/enemy_04.png",
		"textures/enemy_05.png",
		"textures/enemy_06.png",
	)
}

// Setup is called before the main loop starts. It allows you to add entities
// and systems to your Scene.
func (*defaultScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})

	world.AddSystem(&input.System{})
	world.AddSystem(&ui.TextUpdateSystem{})
	world.AddSystem(&ui.ButtonControlSystem{})

	m := navigator.NewMap()
	world.AddSystem(&terminal.System{Map: m})
	world.AddSystem(&opposition.OppositionSystem{Map: m})
	world.AddSystem(&user.UserSystem{Map: m})
}

func main() {
	engo.Run(
		engo.RunOptions{
			AssetsRoot:     "assets",
			FPSLimit:       60,
			Fullscreen:     false,
			Height:         800,
			MSAA:           1,
			NotResizable:   false,
			ScaleOnResize:  true,
			Title:          "Arp 147",
			VSync:          false,
			Width:          1200,
			StandardInputs: true,
		},
		&defaultScene{},
	)
}
