package main

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// TerminalSystem is a scrollable, visual and text input-able system.
type TerminalSystem struct{}

type TerminalViewer struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Remove is called whenever an Entity is removed from the World, in order to
// remove it from this sytem as well.
func (*TerminalSystem) Remove(ecs.BasicEntity) {
	// TODO
}

// Update is ran every frame, with `dt` being the time in seconds since the
// last frame.
func (*TerminalSystem) Update(dt float32) {
	// TODO
}

// New is the initialisation of the System.
func (ts *TerminalSystem) New(w *ecs.World) {
	log.Println("TerminalSystem initialized")

	bkg1 := &TerminalViewer{BasicEntity: ecs.NewBasic()}
	bkg2 := &TerminalViewer{BasicEntity: ecs.NewBasic()}
	bkg3 := &TerminalViewer{BasicEntity: ecs.NewBasic()}

	bkg1.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, 0},
		Width:    800,
		Height:   800,
	}

	tbkg1, err := common.LoadedSprite("textures/bkg_t1.jpg")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	bkg1.RenderComponent = common.RenderComponent{
		Drawable: tbkg1,
		Scale:    engo.Point{1, 1},
	}

	bkg2.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{800, 0},
		Width:    400,
		Height:   400,
	}

	tbkg2, err := common.LoadedSprite("textures/bkg_t2.jpg")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	bkg2.RenderComponent = common.RenderComponent{
		Drawable: tbkg2,
		Scale:    engo.Point{1, 1},
	}

	bkg3.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{800, 400},
		Width:    400,
		Height:   400,
	}

	tbkg3, err := common.LoadedSprite("textures/bkg_t3.jpg")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	bkg3.RenderComponent = common.RenderComponent{
		Drawable: tbkg3,
		Scale:    engo.Point{1, 1},
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&bkg1.BasicEntity, &bkg1.RenderComponent, &bkg1.SpaceComponent)
			sys.Add(&bkg2.BasicEntity, &bkg2.RenderComponent, &bkg2.SpaceComponent)
			sys.Add(&bkg3.BasicEntity, &bkg3.RenderComponent, &bkg3.SpaceComponent)
		}
	}
}
