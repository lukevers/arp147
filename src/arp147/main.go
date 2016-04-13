package main

import (
	"arp147/scenes/mainmenu"
	"engo.io/engo"
	"net/http"
	_ "net/http/pprof"
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

	go http.ListenAndServe(":3030", http.DefaultServeMux)
	engo.Run(options, &mainmenu.Scene{})
}
