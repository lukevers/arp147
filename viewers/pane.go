package viewers

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/lukevers/arp147/ui"
)

type Pane struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	World *ecs.World

	// Entities represent pointers to the basic entity and related
	// render component--this allows us direct access to hiding/showing
	// content in each pane.
	entities map[*ecs.BasicEntity]*common.RenderComponent
}

func NewPane(world *ecs.World) *Pane {
	return &Pane{
		BasicEntity: ecs.NewBasic(),
		World:       world,
		entities:    make(map[*ecs.BasicEntity]*common.RenderComponent),
	}
}

func (p *Pane) RegisterEntity(e *ecs.BasicEntity, r *common.RenderComponent) {
	p.entities[e] = r
}

func (p *Pane) Hide() {
	p.RenderComponent.Hidden = true
	for _, entity := range p.entities {
		entity.Hidden = true
	}
}

func (p *Pane) Show() {
	p.RenderComponent.Hidden = false
	for _, entity := range p.entities {
		entity.Hidden = false
	}
}

func (pane *Pane) AddBackground(point engo.Point) {
	pane.SpaceComponent = common.SpaceComponent{
		Position: point,
		Width:    400,
		Height:   400,
	}

	tbkg, err := common.LoadedSprite("textures/bkg_t2.jpg")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	pane.RenderComponent = common.RenderComponent{
		Drawable: tbkg,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	pane.RenderComponent.SetZIndex(-1)

	for _, system := range pane.World.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&pane.BasicEntity,
				&pane.RenderComponent,
				&pane.SpaceComponent,
			)
		}
	}
}

func RegisterButton(button *ui.Text, viewer Viewer) {
	button.OnClicked(func(entity *ecs.BasicEntity, dt float32) {
		panes := viewer.GetPanes()
		if viewer.GetActiveTab() == button {
			return
		}

		for _, tab := range viewer.GetTabs() {
			tab.Font.FG = color.Alpha16{0x666F}
		}

		for _, pane := range panes {
			pane.Hide()
		}

		button.Font.FG = color.White

		viewer.SetActiveTab(button)
		viewer.SetActivePane(panes[button.Text])
		viewer.GetActivePane().Show()
	}).OnEnter(func(entity *ecs.BasicEntity, dt float32) {
		if viewer.GetActiveTab() == button {
			return
		}

		button.Font.FG = color.Alpha16{0xAAAF}
	}).OnLeave(func(entity *ecs.BasicEntity, dt float32) {
		if viewer.GetActiveTab() == button {
			return
		}

		button.Font.FG = color.Alpha16{0x666F}
	})
}
