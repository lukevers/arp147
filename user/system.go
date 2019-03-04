package user

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type UserSystem struct {
	// .
}

type UserViewer struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Remove is called whenever an Entity is removed from the World, in order to
// remove it from this sytem as well.
func (*UserSystem) Remove(ecs.BasicEntity) {
	// TODO
}

// Update is ran every frame, with `dt` being the time in seconds since the
// last frame.
func (*UserSystem) Update(dt float32) {
	// TODO
}

// New is the initialisation of the System.
func (us *UserSystem) New(w *ecs.World) {
	us.addBackground(w)
	log.Println("UserSystem initialized")
}

func (us *UserSystem) addBackground(w *ecs.World) {
	bkg := &UserViewer{BasicEntity: ecs.NewBasic()}

	bkg.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 800, Y: 400},
		Width:    400,
		Height:   400,
	}

	tbkg, err := common.LoadedSprite("textures/bkg_t3.jpg")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	bkg.RenderComponent = common.RenderComponent{
		Drawable: tbkg,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&bkg.BasicEntity, &bkg.RenderComponent, &bkg.SpaceComponent)
		}
	}
}
