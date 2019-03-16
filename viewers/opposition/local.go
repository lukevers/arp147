package opposition

import (
	"engo.io/engo"
	"github.com/lukevers/arp147/ui"
	"github.com/lukevers/arp147/viewers"
)

func (os *OppositionSystem) addOpposition(pane *viewers.Pane) {
	if os.Map.Center.Planet != nil {
		p := os.Map.Center.Planet
		p.SetPosition(engo.Point{X: 1000, Y: 200})
		p.AddToWorld(os.world)
		pane.RegisterEntity(&p.BasicEntity, &p.RenderComponent)
		return
	}

	text := ui.NewText("Nothing around.")
	text.SetX(933).SetY(197)
	text.Insert(pane.World)
	pane.RegisterEntity(&text.BasicEntity, &text.RenderComponent)
}
