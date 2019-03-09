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
	ship *ship.Ship

	tab  *ui.Text
	tabs []*ui.Text

	pane  *viewers.Pane
	panes map[string]*viewers.Pane

	world *ecs.World
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
	us.panes = make(map[string]*viewers.Pane)

	us.addButtons()
	us.createPanes()

	log.Println("UserSystem initialized")
}

func (us *UserSystem) SetActiveTab(tab *ui.Text) {
	us.tab = tab
}

func (us *UserSystem) GetActiveTab() *ui.Text {
	return us.tab
}

func (us *UserSystem) GetTabs() []*ui.Text {
	return us.tabs
}

func (us *UserSystem) GetActivePane() *viewers.Pane {
	return us.pane
}

func (us *UserSystem) SetActivePane(pane *viewers.Pane) {
	us.pane = pane
}

func (us *UserSystem) GetPanes() map[string]*viewers.Pane {
	return us.panes
}

func (us *UserSystem) addShip(pane *viewers.Pane) {
	us.ship = ship.New()
	us.ship.SetSpritesheet("textures/usership_1.png", 108, 126)
	us.ship.SetPosition(engo.Point{X: 1000, Y: 600})
	us.ship.AddToWorld(pane.World)
	pane.RegisterEntity(&us.ship.BasicEntity, &us.ship.RenderComponent)

	hudShieldLevel := ui.NewText(us.ship.HudShieldLevel())
	pane.RegisterEntity(&hudShieldLevel.BasicEntity, &hudShieldLevel.RenderComponent)
	hudShieldLevel.SetX(990).SetY(435)
	hudShieldLevel.Updatable = func(t *ui.Text) {
		t.Text = us.ship.HudShieldLevel()
	}

	hudHealthLevel := ui.NewText(us.ship.HudShipHealth())
	pane.RegisterEntity(&hudHealthLevel.BasicEntity, &hudHealthLevel.RenderComponent)
	hudHealthLevel.SetX(990).SetY(451)
	hudHealthLevel.Updatable = func(t *ui.Text) {
		t.Text = us.ship.HudShipHealth()
	}

	hudSolarLevel := ui.NewText(us.ship.HudEnergySolarLevel())
	pane.RegisterEntity(&hudSolarLevel.BasicEntity, &hudSolarLevel.RenderComponent)
	hudSolarLevel.SetX(820).SetY(749)
	hudSolarLevel.Updatable = func(t *ui.Text) {
		t.Text = us.ship.HudEnergySolarLevel()
	}

	hudFuelLevel := ui.NewText(us.ship.HudFuelLevel())
	pane.RegisterEntity(&hudFuelLevel.BasicEntity, &hudFuelLevel.RenderComponent)
	hudFuelLevel.SetX(820).SetY(765)
	hudFuelLevel.Updatable = func(t *ui.Text) {
		t.Text = us.ship.HudFuelLevel()
	}

	hudShieldLevel.Insert(pane.World)
	hudHealthLevel.Insert(pane.World)
	hudSolarLevel.Insert(pane.World)
	hudFuelLevel.Insert(pane.World)
}

func (us *UserSystem) addButtons() {
	tlocal := ui.NewText("SHIP")
	tlocal.Font.Size = 12
	tlocal.SetX(820).SetY(416)
	us.tabs = append(us.tabs, tlocal)
	us.tab = tlocal
	viewers.RegisterButton(tlocal, us)
	tlocal.Insert(us.world)

	tmap := ui.NewText("FLOOR")
	tmap.Font.Size = 12
	tmap.Font.FG = color.Alpha16{0x666F}
	tmap.SetX(870).SetY(416)
	us.tabs = append(us.tabs, tmap)
	viewers.RegisterButton(tmap, us)
	tmap.Insert(us.world)
}

func (us *UserSystem) createPanes() {
	paneship := viewers.NewPane(us.world)
	paneship.AddBackground(engo.Point{X: 800, Y: 400})
	us.pane = paneship
	us.panes["SHIP"] = paneship
	us.addShip(paneship)

	panefloor := viewers.NewPane(us.world)
	panefloor.AddBackground(engo.Point{X: 800, Y: 400})
	us.panes["FLOOR"] = panefloor
	panefloor.Hide()
}

func (us *UserSystem) registerButton(t *ui.Text) {
	t.OnClicked(func(entity *ecs.BasicEntity, dt float32) {
		if us.tab == t {
			return
		}

		for _, tab := range us.tabs {
			tab.Font.FG = color.Alpha16{0x666F}
		}

		for _, pane := range us.panes {
			pane.Hide()
		}

		t.Font.FG = color.White

		us.tab = t
		us.pane = us.panes[t.Text]
		us.pane.Show()
	}).OnEnter(func(entity *ecs.BasicEntity, dt float32) {
		if us.tab == t {
			return
		}

		t.Font.FG = color.Alpha16{0xAAAF}
	}).OnLeave(func(entity *ecs.BasicEntity, dt float32) {
		if us.tab == t {
			return
		}

		t.Font.FG = color.Alpha16{0x666F}
	})
}
