package window

import (
	"arp147/computer/line"
	"arp147/ui/background"
	"engo.io/ecs"
	"engo.io/engo"
)

var (
	padding  = 42
	entities []*ecs.Entity
)

func (w *Window) StartSession() {
	// Set the window session to active
	w.Active = true

	width := engo.Width()
	height := engo.Height()

	background.TileWorld(w.world, background.Background{
		Scale:   engo.Point{1, 1},
		Texture: engo.Files.Image("computer_background.png"),
	})

	// Horizontal texture
	texture := engo.Files.Image("computer_horizontal.png")
	W := texture.Width()
	H := texture.Height()

	var x, y float32 = 0, 0
	for {
		var newx float32 = x
		var newy float32 = y

		entity := ecs.NewEntity("RenderSystem")
		entity.AddComponent(engo.NewRenderComponent(texture, engo.Point{1, 1}, "computer"))
		entity.AddComponent(&engo.SpaceComponent{
			Position: engo.Point{
				X: newx,
				Y: newy,
			},
			Width:  W,
			Height: H,
		})

		entities = append(entities, entity)
		w.world.AddEntity(entity)

		if x > width {
			if y == (height - H) {
				break
			} else {
				x = 0
				y = (height - H)
			}
		} else {
			x += W
		}
	}

	// Vertical texture
	texture = engo.Files.Image("computer_vertical.png")
	W = texture.Width()
	H = texture.Height()

	x, y = 0, 0
	for {
		var newx float32 = x
		var newy float32 = y

		entity := ecs.NewEntity("RenderSystem")
		entity.AddComponent(engo.NewRenderComponent(texture, engo.Point{1, 1}, "computer"))
		entity.AddComponent(&engo.SpaceComponent{
			Position: engo.Point{
				X: newx,
				Y: newy,
			},
			Width:  W,
			Height: H,
		})

		entities = append(entities, entity)
		w.world.AddEntity(entity)

		if y > height {
			if x == (width - W) {
				break
			} else {
				x = (width - W)
				y = 0
			}
		} else {
			y += H
		}
	}

	// Top left corner
	texture = engo.Files.Image("computer_corner_top_left.png")
	entity := ecs.NewEntity("RenderSystem")
	entity.AddComponent(engo.NewRenderComponent(texture, engo.Point{1, 1}, "computer"))
	entity.AddComponent(&engo.SpaceComponent{
		Position: engo.Point{
			X: 0,
			Y: 0,
		},
		Width:  texture.Width(),
		Height: texture.Height(),
	})

	entities = append(entities, entity)
	w.world.AddEntity(entity)

	// Top right corner
	texture = engo.Files.Image("computer_corner_top_right.png")
	entity = ecs.NewEntity("RenderSystem")
	entity.AddComponent(engo.NewRenderComponent(texture, engo.Point{1, 1}, "computer"))
	entity.AddComponent(&engo.SpaceComponent{
		Position: engo.Point{
			X: (width - texture.Width()),
			Y: 0,
		},
		Width:  texture.Width(),
		Height: texture.Height(),
	})

	entities = append(entities, entity)
	w.world.AddEntity(entity)

	// Bottom left corner
	texture = engo.Files.Image("computer_corner_bottom_left.png")
	entity = ecs.NewEntity("RenderSystem")
	entity.AddComponent(engo.NewRenderComponent(texture, engo.Point{1, 1}, "computer"))
	entity.AddComponent(&engo.SpaceComponent{
		Position: engo.Point{
			X: 0,
			Y: (height - texture.Height()),
		},
		Width:  texture.Width(),
		Height: texture.Height(),
	})

	entities = append(entities, entity)
	w.world.AddEntity(entity)

	// Bottom right corner
	texture = engo.Files.Image("computer_corner_bottom_right.png")
	entity = ecs.NewEntity("RenderSystem")
	entity.AddComponent(engo.NewRenderComponent(texture, engo.Point{1, 1}, "computer"))
	entity.AddComponent(&engo.SpaceComponent{
		Position: engo.Point{
			X: (width - texture.Width()),
			Y: (height - texture.Height()),
		},
		Width:  texture.Width(),
		Height: texture.Height(),
	})

	entities = append(entities, entity)
	w.world.AddEntity(entity)
}

func (w *Window) StopSession() {
	// Set the window to not active
	w.Active = false

	// Remove all text lines
	for _, line := range w.Lines {
		for _, char := range line.Text {
			char.Remove(w.world)
		}
	}

	w.Lines = make(map[int]*line.Line)
	w.Line = 0

	// Remove all gui elements
	for _, entity := range entities {
		w.world.RemoveEntity(entity)
	}

	entities = *new([]*ecs.Entity)
}
