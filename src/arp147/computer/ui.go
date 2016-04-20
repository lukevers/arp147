package computer

import (
	"arp147/ui/background"
	"engo.io/ecs"
	"engo.io/engo"
)

var ComputerPadding = 42

func (c *Computer) CreateEntityUI(world *ecs.World) {
	width := engo.Width()
	height := engo.Height()

	background.TileWorld(world, background.Background{
		Scale:   engo.Point{1, 1},
		Texture: engo.Files.Image("computer_background.png"),
	})

	// Horizontal texture
	texture := engo.Files.Image("computer_horizontal.png")
	w := texture.Width()
	h := texture.Height()

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
			Width:  w,
			Height: h,
		})

		world.AddEntity(entity)

		if x > width {
			if y == (height - h) {
				break
			} else {
				x = 0
				y = (height - h)
			}
		} else {
			x += w
		}
	}

	// Vertical texture
	texture = engo.Files.Image("computer_vertical.png")
	w = texture.Width()
	h = texture.Height()

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
			Width:  w,
			Height: h,
		})

		world.AddEntity(entity)

		if y > height {
			if x == (width - w) {
				break
			} else {
				x = (width - w)
				y = 0
			}
		} else {
			y += h
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

	world.AddEntity(entity)

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

	world.AddEntity(entity)

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

	world.AddEntity(entity)

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

	world.AddEntity(entity)
}
