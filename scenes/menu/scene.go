package menu

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/lukevers/arp147/clock"
	"github.com/lukevers/arp147/input"
	"github.com/lukevers/arp147/player"
	"github.com/lukevers/arp147/scenes/terminal"
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
	p := player.New()

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(&clock.ClockSystem{Clock: p.Clock})
	world.AddSystem(&input.InputSystem{})
	world.AddSystem(&ui.ButtonControlSystem{})
	world.AddSystem(&ui.LabelUpdateSystem{})

	ui.NewLabel(ui.Label{
		FgColor:   color.White,
		Font:      ui.FontPrimary,
		Size:      16,
		Text:      p.Clock.String(),
		Updatable: p.Clock.String,
		Position: ui.Position{
			Point:    engo.Point{10, 10},
			Position: ui.PositionTopLeft,
		},
	}).AddToWorld(world)

	input.RegisterButtons([]input.Key{
		input.Key{
			Name: "??",
			Keys: []engo.Key{engo.C},
			JustReleased: func() {
				engo.SetScene(&terminal.Scene{Player: p}, false)
			},
		},
	})

	labelTerm := ui.NewLabel(ui.Label{
		FgColor: color.White,
		Font:    ui.FontPrimary,
		Size:    16,
		Text:    "Terminal",
		Position: ui.Position{
			Point:    engo.Point{10, 10},
			Position: ui.PositionBottomLeft,
		},
	})

	labelTerm.OnClicked(func(basic *ecs.BasicEntity, dt float32) {
		engo.SetScene(&terminal.Scene{Player: p}, false)
	})

	labelTerm.OnEnter(func(entity *ecs.BasicEntity, dt float32) {
		labelTerm.FgColor = color.Alpha16{0x6666}
		labelTerm.Render()
	}).OnLeave(func(entity *ecs.BasicEntity, dt float32) {
		labelTerm.FgColor = color.White
		labelTerm.Render()
	})

	labelTerm.AddToWorld(world)
}
