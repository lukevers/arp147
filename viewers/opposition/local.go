package opposition

import (
	"github.com/lukevers/arp147/ui"
	"github.com/lukevers/arp147/viewers"
)

func (os *OppositionSystem) addDefaultOpposition(pane *viewers.Pane) {
	text := ui.NewText("Nothing around.") // 240
	text.SetX(933).SetY(197)
	text.Insert(pane.World)
	pane.RegisterEntity(&text.BasicEntity, &text.RenderComponent)
}
