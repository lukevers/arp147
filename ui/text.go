package ui

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// Text represents a text label that can be added to the world.
type Text struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent

	Font *common.Font
	Text string

	X float32
	Y float32
}

// NewText creates a text based on the string given with all of the defaults.
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

func (t *Text) render() *Text {
	t.RenderComponent.Drawable = common.Text{
		Font: t.Font,
		Text: t.Text,
	}

	t.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: t.X, Y: t.Y},
	}

	return t
}

// Insert adds a text entity to the world.
func (t *Text) Insert(world *ecs.World) {
	t.render()

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
		}
	}
}

// Remove removes the text entity from the world.
func (t *Text) Remove(world *ecs.World) {
	world.RemoveEntity(t.BasicEntity)
}
