package mainmenu

import (
	"github.com/engoengine/engo"
)

func (s *Scene) Preload() {
	engo.Files.AddFromDir("resources/scenes/mainmenu", true)
}
