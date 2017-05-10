package scenes

import (
	"arp147/clock"
	"arp147/logging"
	"arp147/scenes/sandbox"
	"arp147/ships"
	"arp147/ui"
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"
)

type DefaultScene struct {
	ship *ships.TheGerschkin
}

func (*DefaultScene) Type() string {
	return "DefaultScene"
}

func (scene *DefaultScene) Preload() {
	// Load files needed for the scene. If we can't load the files, everything
	// will be fucked up and we should just stop.
	if err := engo.Files.Load(
		"textures/space.png",
		"fonts/CasaleTwo.ttf",
	); err != nil {
		logging.Stderr.Fatal("Could not preload files: ", err)
	}

	// Create a new Gerschkin ship for this scene.
	scene.ship = ships.NewGerschkin()
	scene.ship.Position = ui.Position{
		engo.Point{0, 0},
		ui.PositionCenterCenter,
	}
}

func (scene *DefaultScene) Setup(world *ecs.World) {
	// Add systems needed for the scene.
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(&ui.ButtonControlSystem{})

	// Tile the world background.
	ui.TileWorld(world, "textures/space.png")

	// Create a label for the title
	title := ui.NewLabel(ui.Label{
		Text:     "Arp 147",
		Position: ui.Position{engo.Point{25, 25}, ui.PositionTopLeft},
		Font:     ui.PrimaryFont,
		FgColor:  color.White,
		BgColor:  color.Transparent,
		Size:     64,
	})

	title.OnEnter(func(entity *ecs.BasicEntity, dt float32) {
		title.FgColor = color.Black
		title.BgColor = color.White
		title.Render()
	}).OnLeave(func(entity *ecs.BasicEntity, dt float32) {
		title.FgColor = color.White
		title.BgColor = color.Transparent
		title.Render()
	})

	title.AddToWorld(world)

	// Create a label for quitting the game
	quit := ui.NewLabel(ui.Label{
		Text:     "Quit",
		Position: ui.Position{engo.Point{50, 50}, ui.PositionBottomRight},
		Font:     ui.PrimaryFont,
		FgColor:  color.White,
		BgColor:  color.Transparent,
		Size:     25,
	})

	quit.OnEnter(func(entity *ecs.BasicEntity, dt float32) {
		engo.SetCursor(engo.CursorHand)
	}).OnLeave(func(entity *ecs.BasicEntity, dt float32) {
		engo.SetCursor(engo.CursorNone)
	}).OnClicked(func(entity *ecs.BasicEntity, dt float32) {
		engo.Exit()
	})

	quit.AddToWorld(world)

	// Create a label for starting a new game
	start := ui.NewLabel(ui.Label{
		Text:     "New",
		Position: ui.Position{engo.Point{50, 100}, ui.PositionBottomRight},
		Font:     ui.PrimaryFont,
		FgColor:  color.White,
		BgColor:  color.Transparent,
		Size:     25,
	})

	start.OnEnter(func(entity *ecs.BasicEntity, dt float32) {
		engo.SetCursor(engo.CursorHand)
	}).OnLeave(func(entity *ecs.BasicEntity, dt float32) {
		engo.SetCursor(engo.CursorNone)
	}).OnClicked(func(entity *ecs.BasicEntity, dt float32) {
		engo.SetCursor(engo.CursorNone)
		clock.New()
		engo.SetScene(&sandbox.SandboxScene{}, true)

	})

	start.AddToWorld(world)

	// Add our ship to the world
	scene.ship.AddToWorld(world)
}
