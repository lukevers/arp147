package text

import (
	"arp147/ui/fonts"
	"github.com/engoengine/ecs"
	"github.com/engoengine/engo"
)

type Text struct {
	Text     string
	Size     float64
	Position engo.Point
	Font     fonts.Font
	Scale    engo.Point
	Color    Color

	control *TextControlSystem
	entity  *ecs.Entity
	render  *engo.RenderComponent
	space   *engo.SpaceComponent
}

// Create a new *Text
func New(t Text) *Text {
	t.control = &TextControlSystem{}
	return &t
}

// Update the entity's components that are rendered. The entity replaces each
// component on update, so the GC will handle the deletion of the unused
// components.
func (t *Text) Render() {
	font := &engo.Font{
		Size: t.Size,
		BG:   t.Color.BG,
		FG:   t.Color.FG,
		TTF:  fonts.Get(t.Font),
	}

	t.render = engo.NewRenderComponent(font.Render(t.Text), t.Scale, "text")
	t.entity.AddComponent(t.render)

	x, y, _ := font.TextDimensions(t.Text)
	t.entity.AddComponent(&engo.SpaceComponent{
		Position: t.Position,
		Width:    float32(x),
		Height:   float32(y),
	})
}

// Convert *Text to a usable *ecs.Entity
func (t *Text) Entity(w *ecs.World) *ecs.Entity {
	w.AddSystem(t.control)

	t.entity = ecs.NewEntity([]string{
		"RenderSystem",
		"MouseSystem",
		"TextControlSystem",
	})

	t.Render()
	t.entity.AddComponent(&engo.MouseComponent{})

	return t.entity
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
