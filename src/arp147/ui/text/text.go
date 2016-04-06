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

// Create a new *Text
func New(t Text) *Text {
	t.control = &TextControlSystem{}
	return &t
}

// Convert *Text to a usable *ecs.Entity
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

// Callback function for when the mouse clicked text
func (t *Text) OnClicked(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Clicked = fn
}

// Callback function for when the mouse released text
func (t *Text) OnReleased(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Released = fn
}

// Callback function for when the mouse hovered text
func (t *Text) OnHovered(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Hovered = fn
}

// Callback function for when the mouse dragged text
func (t *Text) OnDragged(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Dragged = fn
}

// Callback function for when the mouse right clicked text
func (t *Text) OnRightClicked(fn func(entity *ecs.Entity, dt float32)) {
	t.control.RightClicked = fn
}

// Callback function for when the mouse right released text
func (t *Text) OnRightReleased(fn func(entity *ecs.Entity, dt float32)) {
	t.control.RightReleased = fn
}

// Callback function for when the mouse enters text
func (t *Text) OnEnter(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Enter = fn
}

// Callback function for when the mouse leaves text
func (t *Text) OnLeave(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Leave = fn
}
