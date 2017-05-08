package scenes

import (
	"arp147/logging"
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
}

func (scene *DefaultScene) Setup(world *ecs.World) {
	// Add systems needed for the scene.
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(&ui.ButtonControlSystem{})

	// Tile the world background.
	ui.TileWorld(world, "textures/space.png")

	// Create a label for the title
	title := ui.NewLabel("Arp 147", ui.PrimaryFont, 64)
	title.SetPosition(ui.Position{
		engo.Point{25, 25},
		ui.PositionTopLeft,
	})

	// Change the colors of the title label on enter/leave.
	title.OnEnter(func(entity *ecs.BasicEntity, dt float32) {
		title.SetForegroundColor(color.Black).SetBackgroundColor(color.White)
	}).OnLeave(func(entity *ecs.BasicEntity, dt float32) {
		title.SetForegroundColor(color.White).SetBackgroundColor(color.Transparent)
	})

	title.AddToWorld(world)

	// Create a label for quitting the game
	quit := ui.NewLabel("Quit", ui.PrimaryFont, 25)
	quit.SetPosition(ui.Position{
		engo.Point{50, 50},
		ui.PositionBottomRight,
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
	start := ui.NewLabel("New Game", ui.PrimaryFont, 25)
	start.SetPosition(ui.Position{
		engo.Point{50, 100},
		ui.PositionBottomRight,
	})

	start.OnEnter(func(entity *ecs.BasicEntity, dt float32) {
		engo.SetCursor(engo.CursorHand)
	}).OnLeave(func(entity *ecs.BasicEntity, dt float32) {
		engo.SetCursor(engo.CursorNone)
	})

	start.AddToWorld(world)

	// Add our ship to the world
	scene.ship.AddToWorld(world, engo.Point{200, 200})
}
