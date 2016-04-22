package main

import (
	//"arp147/scenes/mainmenu"
	"arp147/scenes/newgame"
	"engo.io/engo"
)

func main() {
	options := engo.RunOptions{
		Title:         "Arp 147",
		Fullscreen:    false,
		Width:         1280,
		Height:        780,
		VSync:         false,
		ScaleOnResize: true,
		//FPSLimit: 120,
	}

	//engo.Run(options, &mainmenu.Scene{})
	engo.Run(options, &newgame.Scene{})
}
