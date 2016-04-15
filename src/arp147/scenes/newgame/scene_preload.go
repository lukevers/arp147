package newgame

import (
	"engo.io/engo"
)

func (s *Scene) Preload() {
	engo.Files.AddFromDir("resources/scenes/newgame", true)
	engo.Files.AddFromDir("resources/ships/player", true)
	engo.Files.AddFromDir("resources/ui/computer", true)
}
