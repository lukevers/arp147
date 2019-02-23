package ui

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type Text struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent

	Font *common.Font
	Text string

	X float32
	Y float32
}

func NewText(text string) *Text {
	t := &Text{
		BasicEntity: ecs.NewBasic(),
		Text:        text,
		Font: &common.Font{
			URL:  "fonts/Undefined.ttf",
			FG:   color.White,
			BG:   color.Transparent,
			Size: 14,
		},
	}

	t.Font.CreatePreloaded()
	return t
}

func (t *Text) Render() *Text {
	t.RenderComponent.Drawable = common.Text{
		Font: t.Font,
		Text: t.Text,
	}

	t.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: t.X, Y: t.Y},
		// Width:    float32(t.Font.Size),
		// Height:   float32(t.Font.Size),
		// Width:  0,
		// Height: 0,
	}

	return t
}

func (t *Text) Insert(world *ecs.World) {
	t.Render()

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
		}
	}
}

// Remove removes the text from the world.
func (t *Text) Remove(world *ecs.World) {
	world.RemoveEntity(t.BasicEntity)
}
