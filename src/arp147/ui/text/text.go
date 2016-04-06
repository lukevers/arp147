package text

import (
	"arp147/ui/fonts"
	"github.com/engoengine/ecs"
	"github.com/engoengine/engo"
	"image/color"
)

type Text struct {
	Text     string
	Size     float64
	Position engo.Point
	Font     fonts.Font
	BG       color.Color
	FG       color.Color
	Scale    engo.Point

	control *TextControlSystem
}

func New(t Text) *Text {
	t.control = &TextControlSystem{}
	return &t
}

func (t *Text) Entity(w *ecs.World) *ecs.Entity {
	font := &engo.Font{
		Size: t.Size,
		BG:   t.BG,
		FG:   t.FG,
		TTF:  fonts.Get(t.Font),
	}

	w.AddSystem(t.control)

	entity := ecs.NewEntity([]string{
		"RenderSystem",
		"MouseSystem",
		"TextControlSystem",
	})

	texture := font.Render(t.Text)
	render := engo.NewRenderComponent(texture, t.Scale, "text")
	x, y, _ := font.TextDimensions(t.Text)

	entity.AddComponent(render)
	entity.AddComponent(&engo.MouseComponent{})
	entity.AddComponent(&engo.SpaceComponent{
		Position: t.Position,
		Width:    float32(x),
		Height:   float32(y),
	})

	return entity
}

func (t *Text) OnClicked(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Clicked = fn
}

func (t *Text) OnReleased(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Released = fn
}

func (t *Text) OnHovered(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Hovered = fn
}

func (t *Text) OnDragged(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Dragged = fn
}

func (t *Text) OnRightClicked(fn func(entity *ecs.Entity, dt float32)) {
	t.control.RightClicked = fn
}

func (t *Text) OnRightReleased(fn func(entity *ecs.Entity, dt float32)) {
	t.control.RightReleased = fn
}

func (t *Text) OnEnter(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Enter = fn
}

func (t *Text) OnLeave(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Leave = fn
}
