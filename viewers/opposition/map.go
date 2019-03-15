package opposition

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
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
}
