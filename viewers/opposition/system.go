package opposition

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"github.com/lukevers/arp147/navigator"
	"github.com/lukevers/arp147/ui"
	"github.com/lukevers/arp147/viewers"
)

type OppositionSystem struct {
	Map *navigator.Map

	viewer *viewers.View
	world  *ecs.World
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
func (os *OppositionSystem) New(w *ecs.World) {
	os.world = w
	os.viewer = viewers.New()

	os.addButtons()
	os.createPanes()

	log.Println("OppositionSystem initialized")
}

func (os *OppositionSystem) addButtons() {
	tlocal := ui.NewText("LOCAL")
	tlocal.Font.Size = 12
	tlocal.SetX(820).SetY(16)
	tlocal.Insert(os.world)
	os.viewer.RegisterButton(tlocal)
	os.viewer.SetActiveTab(tlocal)

	tmap := ui.NewText("MAP")
	tmap.Font.Size = 12
	tmap.Font.FG = color.Alpha16{0x666F}
	tmap.SetX(870).SetY(16)
	tmap.Insert(os.world)
	os.viewer.RegisterButton(tmap)
}

func (os *OppositionSystem) createPanes() {
	panelocal := viewers.NewPane(os.world)
	panelocal.AddBackground(engo.Point{X: 800, Y: 0})
	// os.addDefaultOpposition(panelocal)
	os.addRandomPlanet(panelocal)
	os.viewer.AddPane("LOCAL", panelocal)
	os.viewer.SetActivePane(panelocal)

	panemap := viewers.NewPane(os.world)
	panemap.AddBackground(engo.Point{X: 800, Y: 0})
	os.viewer.AddPane("MAP", panemap)
	os.addMap(panemap)
	panemap.Hide()
}
