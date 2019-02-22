package main

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
}

func NewText(text string) *Text {
	t := &Text{
		Text: text,
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

func (t *Text) Render() {
	t.RenderComponent.Drawable = common.Text{
		Font: t.Font,
		Text: t.Text,
	}

	t.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 45, Y: 45},
		Width:    200,
		Height:   200,
	}
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
