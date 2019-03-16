package opposition

import (
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

		y := ui.NewText(cell.HudY())
		y.Font.Size = 12
		y.SetX(820 + xoffset).SetY(119 + yoffset)
		y.Insert(pane.World)
		pane.RegisterEntity(&y.BasicEntity, &y.RenderComponent)

		if cell.Planet != nil {
			cell.Planet.Icon.SetPosition(engo.Point{X: 875 + xoffset, Y: 65 + yoffset})
			cell.Planet.Icon.AddToWorld(pane.World)
			pane.RegisterEntity(&cell.Planet.Icon.BasicEntity, &cell.Planet.Icon.RenderComponent)
		}
	}
}
