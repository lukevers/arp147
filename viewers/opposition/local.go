package opposition

import (
	"engo.io/engo"
	"github.com/lukevers/arp147/planet"
	"github.com/lukevers/arp147/ui"
	"github.com/lukevers/arp147/viewers"
)

func (os *OppositionSystem) addDefaultOpposition(pane *viewers.Pane) {
	text := ui.NewText("Nothing around.") // 240
	text.SetX(933).SetY(197)
	text.Insert(pane.World)
	pane.RegisterEntity(&text.BasicEntity, &text.RenderComponent)
}

func (os *OppositionSystem) addRandomPlanet(pane *viewers.Pane) {
	p := planet.New(planet.SizeViewer, planet.TypePlanet, true)
	p.SetPosition(engo.Point{X: 1000, Y: 200})
	p.AddToWorld(os.world)
	pane.RegisterEntity(&p.BasicEntity, &p.RenderComponent)
}
