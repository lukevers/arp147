package text

import (
	"arp147/ui/fonts"
	"arp147/ui/position"
	"engo.io/ecs"
	"engo.io/engo"
)

type Text struct {
	Text     string
	Size     float64
	Font     fonts.Font
	Scale    engo.Point
	Color    Color
	Position position.Position

	entity *ecs.Entity
	mouse  *Mouse
}

// Create a new *Text
func New(t Text) *Text {
	t.mouse = &Mouse{}
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

	render := engo.NewRenderComponent(font.Render(t.Text), t.Scale, "text")
	t.entity.AddComponent(render)

	x, y, _ := font.TextDimensions(t.Text)
	t.entity.AddComponent(&engo.SpaceComponent{
		Position: t.Position.Calculate(float32(x), float32(y)),
		Width:    float32(x),
		Height:   float32(y),
	})
}

// Convert *Text to a usable *ecs.Entity
func (t *Text) Entity() *ecs.Entity {
	t.entity = ecs.NewEntity(
		"RenderSystem",
		"MouseSystem",
		"TextControlSystem",
	)

	t.Render()
	t.entity.AddComponent(&engo.MouseComponent{})
	t.entity.AddComponent(&TextControlComponent{
		Mouse: t.mouse,
	})

	return t.entity
}

// Callback function for when the mouse clicked text
func (t *Text) OnClicked(fn func(entity *ecs.Entity, dt float32)) {
	t.mouse.Clicked = fn
}

// Callback function for when the mouse released text
func (t *Text) OnReleased(fn func(entity *ecs.Entity, dt float32)) {
	t.mouse.Released = fn
}

// Callback function for when the mouse hovered text
func (t *Text) OnHovered(fn func(entity *ecs.Entity, dt float32)) {
	t.mouse.Hovered = fn
}

// Callback function for when the mouse dragged text
func (t *Text) OnDragged(fn func(entity *ecs.Entity, dt float32)) {
	t.mouse.Dragged = fn
}

// Callback function for when the mouse right clicked text
func (t *Text) OnRightClicked(fn func(entity *ecs.Entity, dt float32)) {
	t.mouse.RightClicked = fn
}

// Callback function for when the mouse right released text
func (t *Text) OnRightReleased(fn func(entity *ecs.Entity, dt float32)) {
	t.mouse.RightReleased = fn
}

// Callback function for when the mouse enters text
func (t *Text) OnEnter(fn func(entity *ecs.Entity, dt float32)) {
	t.mouse.Enter = fn
}

// Callback function for when the mouse leaves text
func (t *Text) OnLeave(fn func(entity *ecs.Entity, dt float32)) {
	t.mouse.Leave = fn
}
