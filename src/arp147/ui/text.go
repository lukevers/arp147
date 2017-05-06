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

// Text ...
type Text struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.MouseComponent

	Text string
	font common.Font
}

// NewText ...
func NewText(text, font string, size float64) *Text {
	t := &Text{
		BasicEntity:     ecs.NewBasic(),
		RenderComponent: common.RenderComponent{},
		SpaceComponent:  common.SpaceComponent{},
		MouseComponent:  common.MouseComponent{},

		Text: text,
		font: common.Font{
			URL:  fmt.Sprintf("fonts/%s.ttf", font),
			FG:   color.White,
			Size: size,
		},
	}

	// Create the font.
	if err := t.font.CreatePreloaded(); err != nil {
		logging.Stderr.Fatal("Could not create preloaded font: ", err)
	}

	// Set the shader to a basic
	t.SetShader(common.HUDShader)
	return t
}

// SetForegroundColor allows you to set the foreground color of the Text.
func (text *Text) SetForegroundColor(c color.Color) {
	text.font.FG = c
}

// SetBackgroundColor allows you to set the background color of the Text.
func (text *Text) SetBackgroundColor(c color.Color) {
	text.font.BG = c
}

// SetPosition allows you to set where on the screen it should be.
func (text *Text) SetPosition(position engo.Point) {
	text.SpaceComponent.Position = position
}

// AddToWorld ...
func (text *Text) AddToWorld(world *ecs.World) {
	text.RenderComponent.Drawable = common.Text{
		Font: &text.font,
		Text: text.Text,
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&text.BasicEntity,
				&text.RenderComponent,
				&text.SpaceComponent,
			)
		}
	}
}
