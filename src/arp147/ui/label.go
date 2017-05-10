package ui

import (
	"arp147/logging"
	"engo.io/ecs"
	"engo.io/engo/common"
	"fmt"
	"image/color"
)

const (
	PrimaryFont = "CasaleTwo"
)

// Lavel is a HUD label that can be placed on the screen. It can be altered on
// mouse events with the ButtonControlComponent.
type Label struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.MouseComponent
	ButtonControlComponent

	Text      string
	Position  Position
	Updatable func() string

	Font string
	Size float64
	font common.Font

	FgColor color.Color
	BgColor color.Color
}

// LabelUpdateSystem is an extension of the
type LabelUpdateSystem struct {
	entities []labelEntity
}

type labelEntity struct {
	*Label
}

// NewLabel creates a new label.
func NewLabel(label Label) *Label {
	label.BasicEntity = ecs.NewBasic()
	label.RenderComponent = common.RenderComponent{}
	label.SpaceComponent = common.SpaceComponent{}
	label.MouseComponent = common.MouseComponent{}
	label.ButtonControlComponent = ButtonControlComponent{}

	label.font = common.Font{
		URL:  fmt.Sprintf("fonts/%s.ttf", label.Font),
		FG:   label.FgColor,
		BG:   label.BgColor,
		Size: label.Size,
	}

	// Create the font.
	if err := label.font.CreatePreloaded(); err != nil {
		logging.Stderr.Fatal("Could not create preloaded font: ", err)
	}

	// Set the shader to a basic
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

// AddToWorld ...
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
