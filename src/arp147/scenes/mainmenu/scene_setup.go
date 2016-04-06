package mainmenu

import (
	"arp147/ui/fonts"
	"arp147/ui/text"
	"github.com/engoengine/ecs"
	"github.com/engoengine/engo"
	"image/color"
	"util/log"
)

func (s *Scene) Setup(w *ecs.World) {
	engo.SetBackground(color.Black)

	w.AddSystem(&engo.MouseSystem{})
	w.AddSystem(&engo.RenderSystem{})

	// Add button
	button := text.New(text.Text{
		Text:     "hello, how are you?",
		Size:     32,
		Position: engo.Point{0, 0},
		Font:     fonts.FONT_PRIMARY,
		BG:       color.Transparent,
		FG:       color.White,
		Scale:    engo.Point{1, 1},
	})

	button.OnClicked(func(entity *ecs.Entity, dt float32) {
		log.Stdout.Println(entity)
	})

	w.AddEntity(button.Entity(w))
}
