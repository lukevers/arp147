package ui

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"image/color"
)

// TODO
const (
	FontPrimary = "fonts/Primary.ttf"
)

// TODO
type Label struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.MouseComponent
	ButtonControlComponent

	BgColor   color.Color
	FgColor   color.Color
	Font      string
	Position  Position
	Size      float64
	Text      string
	Updatable func() string

	font common.Font
}

// TODO
type LabelUpdateSystem struct {
	entities []labelEntity
}

type labelEntity struct {
	*Label
}

func NewLabel(label Label) *Label {
	label.BasicEntity = ecs.NewBasic()
	label.RenderComponent = common.RenderComponent{}
	label.SpaceComponent = common.SpaceComponent{}
	label.MouseComponent = common.MouseComponent{}
	label.ButtonControlComponent = ButtonControlComponent{}

	label.font = common.Font{
		URL:  label.Font,
		FG:   label.FgColor,
		BG:   label.BgColor,
		Size: label.Size,
	}

	// Create the font.
	if err := label.font.CreatePreloaded(); err != nil {
		panic(err)
	}

	label.SetShader(common.HUDShader)

	return &label
}

func (label *Label) Render() {
	label.font.BG = label.BgColor
	label.font.FG = label.FgColor

	w, h, _ := label.font.TextDimensions(label.Text)
	label.SpaceComponent.Width = float32(w)
	label.SpaceComponent.Height = float32(h)

	label.SpaceComponent.Position = label.Position.Calculate(
		label.SpaceComponent.Width,
		label.SpaceComponent.Height,
	)

	label.RenderComponent.Drawable = common.Text{
		Font: &label.font,
		Text: label.Text,
	}
}

// AddToWorld adds the label to the world
func (label *Label) AddToWorld(world *ecs.World) {
	label.Render()

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&label.BasicEntity,
				&label.RenderComponent,
				&label.SpaceComponent,
			)
		case *common.MouseSystem:
			sys.Add(
				&label.BasicEntity,
				&label.MouseComponent,
				&label.SpaceComponent,
				&label.RenderComponent,
			)
		case *ButtonControlSystem:
			sys.Add(
				&label.BasicEntity,
				&label.MouseComponent,
				&label.ButtonControlComponent,
			)
		case *LabelUpdateSystem:
			sys.Add(label)
		}
	}
}

// RemoveFromWorld removes the label from the world
func (label *Label) RemoveFromWorld(world *ecs.World) {
	world.RemoveEntity(label.BasicEntity)
}

// Add takes an entity and adds it to the system
func (l *LabelUpdateSystem) Add(label *Label) {
	l.entities = append(l.entities, labelEntity{label})
}

// Remove takes an entity and removes it from the system
func (l *LabelUpdateSystem) Remove(basic ecs.BasicEntity) {
	delete := -1

	for index, e := range l.entities {
		if e.Label.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}

	if delete >= 0 {
		l.entities = append(l.entities[:delete], l.entities[delete+1:]...)
	}
}

// Update is called on each frame when the system is in use.
func (l *LabelUpdateSystem) Update(dt float32) {
	for _, e := range l.entities {
		// Do nothing with this label if it doesn't have an updatable function.
		if e.Label.Updatable == nil {
			continue
		}

		e.Label.Text = e.Label.Updatable()
		e.Label.Render()
	}
}
