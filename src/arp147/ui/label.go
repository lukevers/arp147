package ui

import (
	"arp147/logging"
	"engo.io/ecs"
	"engo.io/engo"
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

	Text string
	font common.Font
}

// NewLabel creates a new label.
func NewLabel(text, font string, size float64) *Label {
	l := &Label{
		BasicEntity:            ecs.NewBasic(),
		RenderComponent:        common.RenderComponent{},
		SpaceComponent:         common.SpaceComponent{},
		MouseComponent:         common.MouseComponent{},
		ButtonControlComponent: ButtonControlComponent{},

		Text: text,
		font: common.Font{
			URL:  fmt.Sprintf("fonts/%s.ttf", font),
			FG:   color.White,
			Size: size,
		},
	}

	// Create the font.
	if err := l.font.CreatePreloaded(); err != nil {
		logging.Stderr.Fatal("Could not create preloaded font: ", err)
	}

	// Set the shader to a basic
	l.SetShader(common.HUDShader)
	return l
}

// SetForegroundColor allows you to set the foreground color of the Label.
func (label *Label) SetForegroundColor(c color.Color) {
	label.font.FG = c
}

// SetBackgroundColor allows you to set the background color of the Label.
func (label *Label) SetBackgroundColor(c color.Color) {
	label.font.BG = c
}

// SetPosition allows you to set where on the screen it should be.
func (label *Label) SetPosition(position engo.Point) {
	label.SpaceComponent.Position = position
}

// AddToWorld ...
func (label *Label) AddToWorld(world *ecs.World) {
	label.RenderComponent.Drawable = common.Text{
		Font: &label.font,
		Text: label.Text,
	}

	w, h, _ := label.font.TextDimensions(label.Text)
	label.SpaceComponent.Width = float32(w)
	label.SpaceComponent.Height = float32(h)

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
		}
	}
}
