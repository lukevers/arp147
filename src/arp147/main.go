package main

import (
	"arp147/scenes/mainmenu"
	"github.com/engoengine/engo"
)

func main() {
	options := engo.RunOptions{
		Title:         "Arp 147",
		Fullscreen:    false,
		Width:         1200,
		Height:        800,
		VSync:         false,
		ScaleOnResize: true,
		//FPSLimit: 120,
	}

	engo.Run(options, &mainmenu.Scene{})
}
