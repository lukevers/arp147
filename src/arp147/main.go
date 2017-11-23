package main

import (
	_ "arp147/flags"
	//	"arp147/logging"
	"arp147/scenes"
	"engo.io/engo"
)

func main() {
	/*
		// Catch all panics and log them.
		defer func() {
			if r := recover(); r != nil {
				logging.Stderr.Fatal(r)
			}
		}()

	*/

	// Setup our run options
	options := engo.RunOptions{
		Title:  "Arp 147",
		Width:  1200,
		Height: 800,
	}

	// Run the game
	engo.Run(options, &scenes.DefaultScene{})
}
