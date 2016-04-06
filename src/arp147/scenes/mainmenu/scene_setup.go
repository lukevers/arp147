package mainmenu

import (
	"arp147/ui/fonts"
	"arp147/ui/position"
	"arp147/ui/text"
	"github.com/engoengine/ecs"
	"github.com/engoengine/engo"
	"image/color"
)

func (s *Scene) Setup(w *ecs.World) {
	engo.SetBackground(color.Black)

	w.AddSystem(&engo.MouseSystem{})
	w.AddSystem(&engo.RenderSystem{})

	title := text.New(text.Text{
		Text:  "Arp 147",
		Size:  64,
		Font:  fonts.FONT_PRIMARY,
		Scale: engo.Point{1, 1},
		Color: text.Color{
			BG: color.Transparent,
			FG: color.White,
		},
		Position: position.Position{
			Point:    engo.Point{25, 25},
			Position: position.BOTTOM_RIGHT,
		},
	})

	title.OnEnter(func(entity *ecs.Entity, dt float32) {
		title.Color = text.Color{
			BG: color.White,
			FG: color.Black,
		}

		title.Render()
	})

	title.OnLeave(func(entity *ecs.Entity, dt float32) {
		title.Color = text.Color{
			BG: color.Transparent,
			FG: color.White,
		}

		title.Render()
	})

	w.AddEntity(title.Entity(w))
}
