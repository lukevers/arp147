package mainmenu

import (
	"arp147/scenes/mainmenu/display"
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
	w.AddSystem(&display.FakePlayerSystem{})

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

	entity := ecs.NewEntity("RenderSystem", "MouseSystem", "FakePlayerSystem")
	texture := engo.Files.Image("ship-00-no-shields.png")
	entity.AddComponent(engo.NewRenderComponent(texture, engo.Point{1, 1}, "player"))
	pos := position.Position{
		Point:    engo.Point{0, 0},
		Position: position.CENTER_CENTER,
	}

	var begin bool
	entity.AddComponent(&engo.MouseComponent{})
	entity.AddComponent(&display.FakePlayerComponent{
		Begin: &begin,
	})

	entity.AddComponent(&engo.SpaceComponent{
		Position: pos.Calculate(float32(texture.Width()), float32(texture.Height())),
		Width:    texture.Width(),
		Height:   texture.Height(),
	})

	w.AddEntity(entity)

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
