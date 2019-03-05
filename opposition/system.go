package opposition

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type OppositionSystem struct {
	// .
}

type OppositionViewer struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Remove is called whenever an Entity is removed from the World, in order to
// remove it from this sytem as well.
func (*OppositionSystem) Remove(ecs.BasicEntity) {
	// TODO
}

// Update is ran every frame, with `dt` being the time in seconds since the
// last frame.
func (*OppositionSystem) Update(dt float32) {
	// TODO
}

// New is the initialisation of the System.
func (us *OppositionSystem) New(w *ecs.World) {
	us.addBackground(w)
	log.Println("OppositionSystem initialized")
}

func (us *OppositionSystem) addBackground(w *ecs.World) {
	bkg := &OppositionViewer{BasicEntity: ecs.NewBasic()}

	bkg.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 800, Y: 0},
		Width:    400,
		Height:   400,
	}

	tbkg, err := common.LoadedSprite("textures/bkg_t2.jpg")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	bkg.RenderComponent = common.RenderComponent{
		Drawable: tbkg,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	bkg.RenderComponent.SetZIndex(0)

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&bkg.BasicEntity, &bkg.RenderComponent, &bkg.SpaceComponent)
		}
	}
}
