package mainmenu

import (
	"arp147/ships"
	"arp147/ships/player/gerschkin"
	"arp147/ui/background"
	"arp147/ui/fonts"
	"arp147/ui/position"
	"arp147/ui/text"
	"engo.io/ecs"
	"engo.io/engo"
	"image/color"
)

func (s *Scene) Setup(w *ecs.World) {
	engo.SetBackground(color.Black)

	w.AddSystem(&engo.MouseSystem{})
	w.AddSystem(&engo.RenderSystem{})
	w.AddSystem(&text.TextControlSystem{})

	// -- Background

	background.TileWorld(w, background.Background{
		Scale:   engo.Point{1, 1},
		Texture: engo.Files.Image("space.png"),
	})

	// -- Title

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

	w.AddEntity(title.Entity())

	// -- Quit Game

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

	w.AddEntity(quit.Entity())

	// -- Fake Player

	begin := false
	ship := ships.New(&gerschkin.Ship{}).(*gerschkin.Ship).Entity("ShipSystem")
	w.AddSystem(&ShipSystem{})
	ship.AddComponent(&ShipComponent{Begin: &begin})
	w.AddEntity(ship)

	// -- New Game

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
			Point:    engo.Point{50, 100},
			Position: position.BOTTOM_RIGHT,
		},
	})

	newGame.OnEnter(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(engo.Hand)
	})

	newGame.OnLeave(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(nil)
	})

	newGame.OnClicked(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(nil)
		begin = true
	})

	w.AddEntity(newGame.Entity())
}
