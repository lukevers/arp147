package background

import (
	"engo.io/ecs"
	"engo.io/engo"
)

type Background struct {
	Scale   engo.Point
	Texture *engo.Texture
}

// TileWorld is a func that takes a texture and tiles it on the entire screen.
// It's a function that takes a world and Background as parameters because this
// is generally going to be a function that's used for tiling backgrounds.
func TileWorld(world *ecs.World, b Background) {
	width := engo.Width()
	height := engo.Height()

	w := b.Texture.Width()
	h := b.Texture.Height()

	var x, y float32 = 0, 0

	for {
		var newx float32 = x
		var newy float32 = y

		entity := ecs.NewEntity("RenderSystem")
		entity.AddComponent(engo.NewRenderComponent(b.Texture, b.Scale, "background"))
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
			x = 0
			y += h
		} else {
			x += w
		}

		if y > height {
			break
		}
	}
}
