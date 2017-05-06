package ui

import (
	"arp147/logging"
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type background struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// TileWorld is a helper function that tiles an image on the entire screen.
func TileWorld(world *ecs.World, image string) {
	texture, err := common.LoadedSprite(image)
	if err != nil {
		logging.Stderr.Fatal("Could not use loaded sprite: ", err)
	}

	width := engo.GameWidth()
	height := engo.GameHeight()

	w := texture.Width()
	h := texture.Height()

	var x, y float32 = 0, 0

	for {
		var x2 float32 = x
		var y2 float32 = y

		bkg := background{
			BasicEntity: ecs.NewBasic(),
			SpaceComponent: common.SpaceComponent{
				Position: engo.Point{x2, y2},
				Width:    w,
				Height:   h,
			},
			RenderComponent: common.RenderComponent{
				Drawable: texture,
				Scale:    engo.Point{1, 1},
			},
		}

		for _, system := range world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(
					&bkg.BasicEntity,
					&bkg.RenderComponent,
					&bkg.SpaceComponent,
				)
			}
		}

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
