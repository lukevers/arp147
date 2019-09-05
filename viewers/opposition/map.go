package opposition

import (
	"fmt"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/lukevers/arp147/ui"
	"github.com/lukevers/arp147/viewers"
)

func (os *OppositionSystem) addMap(pane *viewers.Pane) {
	os.addGrid(pane)
}

func (os *OppositionSystem) addGrid(pane *viewers.Pane) {
	type grid struct {
		ecs.BasicEntity
		common.RenderComponent
		common.SpaceComponent
	}

	g := grid{BasicEntity: ecs.NewBasic()}
	g.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 800, Y: 0},
		Width:    400,
		Height:   400,
	}

	tbkg, err := common.LoadedSprite("textures/bkg_t2_grid.png")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	g.RenderComponent = common.RenderComponent{
		Drawable: tbkg,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	for _, system := range os.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&g.BasicEntity,
				&g.RenderComponent,
				&g.SpaceComponent,
			)
		}
	}

	pane.RegisterEntity(&g.BasicEntity, &g.RenderComponent)
	var entities []*ecs.BasicEntity

	genMap := func() {
		for _, entity := range entities {
			pane.RemoveEntity(entity)
			os.world.RemoveEntity(*entity)
		}

		for i, cell := range os.Map.GetVisibleCells() {
			var xoffset, yoffset float32
			if i%3 != 0 {
				xoffset = float32((i % 3) * 125)
			}

			if i >= 3 && i <= 5 {
				yoffset = 125
			} else if i >= 6 && i <= 9 {
				yoffset = 250
			}

			x := ui.NewText(cell.HudX())
			x.Font.Size = 12
			x.SetX(820 + xoffset).SetY(107 + yoffset)
			x.Insert(pane.World)
			pane.RegisterEntity(&x.BasicEntity, &x.RenderComponent)
			entities = append(entities, &x.BasicEntity)

			y := ui.NewText(cell.HudY())
			y.Font.Size = 12
			y.SetX(820 + xoffset).SetY(119 + yoffset)
			y.Insert(pane.World)
			pane.RegisterEntity(&y.BasicEntity, &y.RenderComponent)
			entities = append(entities, &y.BasicEntity)

			etext := "NIL"
			if !cell.Discovered {
				etext = ""
			} else {
				if cell.Entity != nil {
					if yoffset == 0 {
						yoffset = 10
					}

					cell.Entity.Icon.SetPosition(engo.Point{X: 875 + xoffset, Y: 70 + yoffset})
					cell.Entity.Icon.AddToWorld(pane.World)
					pane.RegisterEntity(&cell.Entity.Icon.BasicEntity, &cell.Entity.Icon.RenderComponent)
					entities = append(entities, &cell.Entity.Icon.BasicEntity)
					etext = cell.Entity.Type()
				}
			}

			var e float32 = 35
			if i >= 3 && i <= 5 {
				e = 143
			} else if i >= 6 && i <= 9 {
				e = 264
			}

			entity := ui.NewText(fmt.Sprintf("E: %s", etext))
			entity.Font.Size = 12
			entity.SetX(820 + xoffset).SetY(e)
			entity.Insert(pane.World)
			pane.RegisterEntity(&entity.BasicEntity, &entity.RenderComponent)
			entities = append(entities, &entity.BasicEntity)
		}

		if pane != os.viewer.GetActivePane() {
			pane.Hide()
		}
	}

	genMap()
	engo.Mailbox.Listen("MoveMessage", func(msg engo.Message) {
		genMap()
	})
}
