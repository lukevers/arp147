package opposition

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/lukevers/arp147/ui"
	"github.com/lukevers/arp147/viewers"
)

func (os *OppositionSystem) addOpposition(pane *viewers.Pane) {
	var entities []*ecs.BasicEntity

	genLocal := func() {
		for _, entity := range entities {
			pane.RemoveEntity(entity)
			os.world.RemoveEntity(*entity)
		}

		if os.Map.Center.Entity != nil {
			e := os.Map.Center.Entity
			e.SetPosition(engo.Point{X: 1000, Y: 200})
			e.AddToWorld(os.world)
			pane.RegisterEntity(&e.BasicEntity, &e.RenderComponent)
			entities = append(entities, &e.BasicEntity)

			if pane != os.viewer.GetActivePane() {
				pane.Hide()
			}

			return
		}

		text := ui.NewText("Nothing around.")
		text.SetX(933).SetY(197)
		text.Insert(pane.World)
		pane.RegisterEntity(&text.BasicEntity, &text.RenderComponent)
		entities = append(entities, &text.BasicEntity)

		if pane != os.viewer.GetActivePane() {
			pane.Hide()
		}
	}

	genLocal()
	engo.Mailbox.Listen("MoveMessage", func(msg engo.Message) {
		genLocal()
	})
}
