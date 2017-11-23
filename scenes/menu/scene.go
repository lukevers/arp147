package menu

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/lukevers/arp147/clock"
	"github.com/lukevers/arp147/ui"
	"image/color"
)

// Scene defines a scene for the main menu
type Scene struct{}

// Type defines the scene
func (s *Scene) Type() string {
	return "MenuScene"
}

// Preload
func (s *Scene) Preload() {
	if err := engo.Files.Load(
		ui.FontPrimary,
	); err != nil {
		panic(err)
	}
}

// Setup
func (s *Scene) Setup(world *ecs.World) {
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&clock.ClockSystem{})
	world.AddSystem(&ui.LabelUpdateSystem{})

	clock.New()
	ui.NewLabel(ui.Label{
		FgColor:   color.White,
		Font:      ui.FontPrimary,
		Size:      16,
		Text:      clock.String(),
		Updatable: clock.String,
		Position: ui.Position{
			Point:    engo.Point{10, 10},
			Position: ui.PositionTopLeft,
		},
	}).AddToWorld(world)
}
