package main

import (
	_ "arp147/flags"
	_ "arp147/logging"
	"arp147/scenes"
	"engo.io/engo"
)

func main() {
	options := engo.RunOptions{
		Title:  "Arp 147",
		Width:  1200,
		Height: 800,
	}

	engo.Run(options, &scenes.DefaultScene{})
}
