package mainmenu

import (
	"engo.io/engo"
)

func (s *Scene) Preload() {
	engo.Files.AddFromDir("resources/scenes/mainmenu", true)
	engo.Files.AddFromDir("resources/ships/player/gerschkin", true)
}
