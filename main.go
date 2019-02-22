package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
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
	)
}

// Setup is called before the main loop starts. It allows you to add entities
// and systems to your Scene.
func (*defaultScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})

	world.AddSystem(&TerminalSystem{})

	t := NewText("arp-sh1$")
	t.Render()
	t.Insert(world)
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
