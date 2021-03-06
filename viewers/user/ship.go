package user

import (
	"engo.io/engo"
	"github.com/lukevers/arp147/ship"
	"github.com/lukevers/arp147/ship/user"
	"github.com/lukevers/arp147/ui"
	"github.com/lukevers/arp147/viewers"
)

func (us *UserSystem) addShip(pane *viewers.Pane) {
	us.ship = ship.New(user.User01{})
	us.ship.SetPosition(engo.Point{X: 1000, Y: 600})
	us.ship.AddToWorld(pane.World)
	pane.RegisterEntity(&us.ship.BasicEntity, &us.ship.RenderComponent)

	hudX := ui.NewText(us.Map.Center.HudX())
	pane.RegisterEntity(&hudX.BasicEntity, &hudX.RenderComponent)
	hudX.SetX(820).SetY(435)
	hudX.Updatable = func(t *ui.Text) {
		t.Text = us.Map.Center.HudX()
	}

	hudY := ui.NewText(us.Map.Center.HudY())
	pane.RegisterEntity(&hudY.BasicEntity, &hudY.RenderComponent)
	hudY.SetX(820).SetY(451)
	hudY.Updatable = func(t *ui.Text) {
		t.Text = us.Map.Center.HudY()
	}

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

	hudX.Insert(pane.World)
	hudY.Insert(pane.World)
	hudShieldLevel.Insert(pane.World)
	hudHealthLevel.Insert(pane.World)
	hudSolarLevel.Insert(pane.World)
	hudFuelLevel.Insert(pane.World)
}
