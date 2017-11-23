package main

import (
	"engo.io/engo"
	"github.com/lukevers/arp147/scenes/menu"
)

func main() {
	options := engo.RunOptions{
		AssetsRoot:    "assets",
		FPSLimit:      60,
		Fullscreen:    false,
		Height:        800,
		MSAA:          1,
		NotResizable:  true,
		ScaleOnResize: false,
		Title:         "Arp 147",
		VSync:         false,
		Width:         1200,
	}

	engo.Run(options, &menu.Scene{})
}
