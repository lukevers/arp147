package user

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"github.com/lukevers/arp147/ship"
	"github.com/lukevers/arp147/ui"
	"github.com/lukevers/arp147/viewers"
)

type UserSystem struct {
	ship   *ship.Ship
	viewer *viewers.View
	world  *ecs.World
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
	us.world = w
	us.viewer = viewers.New()

	us.addButtons()
	us.createPanes()

	log.Println("UserSystem initialized")
}

func (us *UserSystem) addButtons() {
	tlocal := ui.NewText("SHIP")
	tlocal.Font.Size = 12
	tlocal.SetX(820).SetY(416)
	us.viewer.RegisterButton(tlocal)
	us.viewer.SetActiveTab(tlocal)
	tlocal.Insert(us.world)

	tmap := ui.NewText("FLOOR")
	tmap.Font.Size = 12
	tmap.Font.FG = color.Alpha16{0x666F}
	tmap.SetX(865).SetY(416)
	us.viewer.RegisterButton(tmap)
	tmap.Insert(us.world)
}

func (us *UserSystem) createPanes() {
	paneship := viewers.NewPane(us.world)
	paneship.AddBackground(engo.Point{X: 800, Y: 400})
	us.viewer.AddPane("SHIP", paneship)
	us.viewer.SetActivePane(paneship)
	us.addShip(paneship)

	panefloor := viewers.NewPane(us.world)
	panefloor.AddBackground(engo.Point{X: 800, Y: 400})
	us.viewer.AddPane("FLOOR", panefloor)
	panefloor.Hide()
}
