package player

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/lukevers/arp147/ui"
	"log"
)

// TODO
type Line struct {
	// TODO
	Text []*ui.Label

	// TODO
	Locked bool
}

// TODO
type Buffer struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	// TODO
	Lines map[int]*Line

	// TODO
	Line int
}

// TODO
type Terminal struct {
	Buffers []*Buffer
}

// TODO
func NewTerminal() *Terminal {
	return &Terminal{}
}

// TODO
func (t *Terminal) Show(world *ecs.World) {
	log.Println("x")

	t.drawUi(world)
}

func (t *Terminal) drawUi(world *ecs.World) {
	w := engo.GameWidth()
	h := engo.GameHeight()

	var buffers []*Buffer
	var x, y float32 = 25, 0

	// Draw the horizontal borders
	for {
		buffer := Buffer{BasicEntity: ecs.NewBasic()}
		buffer.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{x, y},
			Width:    25,
			Height:   25,
		}

		texture, err := common.LoadedSprite("textures/terminal/horizontal.png")
		if err != nil {
			panic(err)
		}

		buffer.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{1, 1},
		}

		buffers = append(buffers, &buffer)

		if x >= (w - 50) {
			if y == (h - 25) {
				break
			} else {
				y = h - 25
				x = 25
			}
		} else {
			x += 25
		}
	}

	// Draw the vertical borders
	x, y = 0, 25
	for {
		buffer := Buffer{BasicEntity: ecs.NewBasic()}
		buffer.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{x, y},
			Width:    25,
			Height:   25,
		}

		texture, err := common.LoadedSprite("textures/terminal/vertical.png")
		if err != nil {
			panic(err)
		}

		buffer.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{1, 1},
		}

		buffers = append(buffers, &buffer)

		if y >= (h - 50) {
			if x == (w - 25) {
				break
			} else {
				x = w - 25
				y = 25
			}
		} else {
			y += 25
		}
	}

	// Draw the corners
	corners := []string{
		"corner_top_left",
		"corner_top_right",
		"corner_bottom_left",
		"corner_bottom_right",
	}

	positions := []engo.Point{
		engo.Point{0, 0},
		engo.Point{w - 28, 0},
		engo.Point{0, h - 28},
		engo.Point{w - 28, h - 28},
	}

	for i := 0; i < 4; i++ {
		buffer := Buffer{BasicEntity: ecs.NewBasic()}
		buffer.SpaceComponent = common.SpaceComponent{
			Position: positions[i],
			Width:    28,
			Height:   28,
		}

		texture, err := common.LoadedSprite("textures/terminal/" + corners[i] + ".png")
		if err != nil {
			panic(err)
		}

		buffer.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{1, 1},
		}

		buffers = append(buffers, &buffer)
	}

	// Draw everything
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, buffer := range buffers {
				sys.Add(
					&buffer.BasicEntity,
					&buffer.RenderComponent,
					&buffer.SpaceComponent,
				)
			}
		}
	}
}
