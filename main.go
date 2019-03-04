package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/lukevers/arp147/input"
	"github.com/lukevers/arp147/opposition"
	"github.com/lukevers/arp147/terminal"
	"github.com/lukevers/arp147/ui"
	"github.com/lukevers/arp147/user"
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
		"textures/bkg_t3.jpg",

		"textures/usership_1.png",
	)
}

// Setup is called before the main loop starts. It allows you to add entities
// and systems to your Scene.
func (*defaultScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&input.InputSystem{})
	world.AddSystem(&ui.TextUpdateSystem{})

	world.AddSystem(&terminal.TerminalSystem{})
	world.AddSystem(&opposition.OppositionSystem{})
	world.AddSystem(&user.UserSystem{})
}

func main() {
	opts := engo.RunOptions{
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
	}
	engo.Run(opts, &defaultScene{})
}
