package ui

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// Text represents text that can be added to the world.
type Text struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	common.MouseComponent
	ButtonControlComponent

	Font      *common.Font
	Text      string
	Updatable func(*Text)

	X float32
	Y float32
}

// NewText creates a text based on the string given with all of the defaults.
func NewText(text string) *Text {
	t := &Text{
		BasicEntity:            ecs.NewBasic(),
		ButtonControlComponent: ButtonControlComponent{},
		MouseComponent:         common.MouseComponent{},
		Text:                   text,
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

// SetY changes the Y position of the text.
func (t *Text) SetY(y float32) *Text {
	t.Y = y
	t.SpaceComponent.Position.Y = y

	return t
}

// SetX changes the X position of the text.
func (t *Text) SetX(x float32) *Text {
	t.X = x
	t.SpaceComponent.Position.X = x

	return t
}

// Render takes the text and displays it to the world it exists in.
func (t *Text) Render() *Text {
	t.RenderComponent.Drawable = common.Text{
		Font: t.Font,
		Text: t.Text,
	}

	t.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: t.X, Y: t.Y},
		Width:    float32(t.Font.Size) * float32(len(t.Text)),
		Height:   float32(t.Font.Size),
	}

	return t
}

// Insert adds a text entity to the world.
func (t *Text) Insert(world *ecs.World) {
	t.Render()

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&t.BasicEntity,
				&t.RenderComponent,
				&t.SpaceComponent,
			)
		case *common.MouseSystem:
			sys.Add(
				&t.BasicEntity,
				&t.MouseComponent,
				&t.SpaceComponent,
				&t.RenderComponent,
			)
		case *TextUpdateSystem:
			sys.Add(t)
		case *ButtonControlSystem:
			sys.Add(
				&t.BasicEntity,
				&t.MouseComponent,
				&t.ButtonControlComponent,
			)
		}
	}
}

// Remove removes the text entity from the world.
func (t *Text) Remove(world *ecs.World) {
	world.RemoveEntity(t.BasicEntity)
}

// TextUpdateSystem allows text to be dynamically updated.
type TextUpdateSystem struct {
	entities []textEntity
}

type textEntity struct {
	*Text
}

// Add takes an entity and adds it to the system
func (t *TextUpdateSystem) Add(text *Text) {
	t.entities = append(t.entities, textEntity{text})
}

// Remove takes an entity and removes it from the system
func (t *TextUpdateSystem) Remove(basic ecs.BasicEntity) {
	delete := -1

	for index, e := range t.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}

	if delete >= 0 {
		t.entities = append(t.entities[:delete], t.entities[delete+1:]...)
	}
}

// Update is called on each frame when the system is in use.
func (t *TextUpdateSystem) Update(dt float32) {
	for _, e := range t.entities {
		// Do nothing with this text if it doesn't have an updatable function.
		if e.Updatable == nil {
			continue
		}

		e.Updatable(e.Text)
		e.Render()
	}
}
