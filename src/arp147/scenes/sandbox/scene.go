package sandbox

import (
	"arp147/clock"
	"arp147/input"
	"arp147/logging"
	"arp147/ui"
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"
)

type SandboxScene struct {
	//
}

func (*SandboxScene) Type() string {
	return "SandboxScene"
}

func (scene *SandboxScene) Preload() {
	// Load files needed for the scene. If we can't load the files, everything
	// will be fucked up and we should just stop.
	if err := engo.Files.Load(
		"textures/space.png",
	); err != nil {
		logging.Stderr.Fatal("Could not preload files: ", err)
	}
}

func (scene *SandboxScene) Setup(world *ecs.World) {
	// Add systems needed for the scene.
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&input.InputSystem{})
	world.AddSystem(&clock.ClockSystem{})
	world.AddSystem(&ui.LabelUpdateSystem{})

	// Tile the world background.
	ui.TileWorld(world, "textures/space.png")

	// Setup supported input for the scene
	scene.SetupInput()

	// Add HUD
	ui.NewLabel(ui.Label{
		Text:      clock.String(),
		Updatable: clock.String,
		Position:  ui.Position{engo.Point{10, 10}, ui.PositionTopLeft},
		Font:      ui.PrimaryFont,
		FgColor:   color.White,
		BgColor:   color.Transparent,
		Size:      16,
	}).AddToWorld(world)
}
