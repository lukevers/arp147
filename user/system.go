package user

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/lukevers/arp147/ship"
	"github.com/lukevers/arp147/ui"
)

type UserSystem struct {
	ship *ship.Ship
}

type UserViewer struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
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
	us.addBackground(w)
	us.addShip(w)
	log.Println("UserSystem initialized")
}

func (us *UserSystem) addBackground(w *ecs.World) {
	bkg := &UserViewer{BasicEntity: ecs.NewBasic()}

	bkg.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 800, Y: 400},
		Width:    400,
		Height:   400,
	}

	tbkg, err := common.LoadedSprite("textures/bkg_t3.jpg")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	bkg.RenderComponent = common.RenderComponent{
		Drawable: tbkg,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&bkg.BasicEntity, &bkg.RenderComponent, &bkg.SpaceComponent)
		}
	}
}

func (us *UserSystem) addShip(w *ecs.World) {
	us.ship = ship.New()
	us.ship.SetSpritesheet("textures/usership_1.png", 108, 126)
	us.ship.SetPosition(engo.Point{X: 1000, Y: 600})
	us.ship.AddToWorld(w)

	hudShieldLevel := ui.NewText(us.ship.HudShieldLevel())
	hudShieldLevel.SetX(990).SetY(420)
	hudShieldLevel.Updatable = func(t *ui.Text) {
		t.Text = us.ship.HudShieldLevel()
	}

	hudHealthLevel := ui.NewText(us.ship.HudShipHealth())
	hudHealthLevel.SetX(990).SetY(436)
	hudHealthLevel.Updatable = func(t *ui.Text) {
		t.Text = us.ship.HudShipHealth()
	}

	hudSolarLevel := ui.NewText(us.ship.HudEnergySolarLevel())
	hudSolarLevel.SetX(820).SetY(749)
	hudSolarLevel.Updatable = func(t *ui.Text) {
		t.Text = us.ship.HudEnergySolarLevel()
	}

	hudFuelLevel := ui.NewText(us.ship.HudFuelLevel())
	hudFuelLevel.SetX(820).SetY(765)
	hudFuelLevel.Updatable = func(t *ui.Text) {
		t.Text = us.ship.HudFuelLevel()
	}

	hudShieldLevel.Insert(w)
	hudHealthLevel.Insert(w)
	hudSolarLevel.Insert(w)
	hudFuelLevel.Insert(w)
}
