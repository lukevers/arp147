package mainmenu

import (
	"arp147/ui/fonts"
	"arp147/ui/position"
	"arp147/ui/text"
	"engo.io/ecs"
	"engo.io/engo"
	"image/color"
	"util/log"
)

func (s *Scene) Setup(w *ecs.World) {
	engo.SetBackground(color.Black)

	w.AddSystem(&engo.MouseSystem{})
	w.AddSystem(&engo.RenderSystem{})
	w.AddSystem(&text.TextControlSystem{})

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
			Position: position.TOP_LEFT,
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

	// ---

	quit := text.New(text.Text{
		Text:  "Quit",
		Size:  25,
		Font:  fonts.FONT_PRIMARY,
		Scale: engo.Point{1, 1},
		Color: text.Color{
			BG: color.Transparent,
			FG: color.White,
		},
		Position: position.Position{
			Point:    engo.Point{50, 50},
			Position: position.BOTTOM_RIGHT,
		},
	})

	quit.OnEnter(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(engo.Hand)
	})

	quit.OnLeave(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(nil)
	})

	quit.OnClicked(func(entity *ecs.Entity, dt float32) {
		engo.Exit()
	})

	w.AddEntity(quit.Entity(w))

	settings := text.New(text.Text{
		Text:  "Settings",
		Size:  25,
		Font:  fonts.FONT_PRIMARY,
		Scale: engo.Point{1, 1},
		Color: text.Color{
			BG: color.Transparent,
			FG: color.White,
		},
		Position: position.Position{
			Point:    engo.Point{50, 100},
			Position: position.BOTTOM_RIGHT,
		},
	})

	settings.OnEnter(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(engo.Hand)
		log.Stdout.Println("y")
	})

	settings.OnLeave(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(nil)
	})

	w.AddEntity(settings.Entity(w))

	newGame := text.New(text.Text{
		Text:  "New Game",
		Size:  25,
		Font:  fonts.FONT_PRIMARY,
		Scale: engo.Point{1, 1},
		Color: text.Color{
			BG: color.Transparent,
			FG: color.White,
		},
		Position: position.Position{
			Point:    engo.Point{50, 150},
			Position: position.BOTTOM_RIGHT,
		},
	})

	newGame.OnEnter(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(engo.Hand)
		log.Stdout.Println("x")
	})

	newGame.OnLeave(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(nil)
	})

	w.AddEntity(newGame.Entity(w))
}
