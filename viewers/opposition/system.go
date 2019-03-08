package opposition

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"github.com/lukevers/arp147/ui"
	"github.com/lukevers/arp147/viewers"
)

type OppositionSystem struct {
	tab  *ui.Text
	tabs []*ui.Text

	pane  *viewers.Pane
	panes map[string]*viewers.Pane

	world *ecs.World
}

// Remove is called whenever an Entity is removed from the World, in order to
// remove it from this sytem as well.
func (*OppositionSystem) Remove(ecs.BasicEntity) {
	// TODO
}

// Update is ran every frame, with `dt` being the time in seconds since the
// last frame.
func (*OppositionSystem) Update(dt float32) {
	// TODO
}

// New is the initialisation of the System.
func (us *OppositionSystem) New(w *ecs.World) {
	us.world = w
	us.panes = make(map[string]*viewers.Pane)

	us.addButtons()
	us.createPanes()

	log.Println("OppositionSystem initialized")
}

func (us *OppositionSystem) addButtons() {
	tlocal := ui.NewText("LOCAL")
	tlocal.Font.Size = 12
	tlocal.SetX(820).SetY(16)
	us.tabs = append(us.tabs, tlocal)
	us.tab = tlocal
	us.registerButton(tlocal)
	tlocal.Insert(us.world)

	tmap := ui.NewText("MAP")
	tmap.Font.Size = 12
	tmap.Font.FG = color.Alpha16{0x666F}
	tmap.SetX(870).SetY(16)
	us.tabs = append(us.tabs, tmap)
	us.registerButton(tmap)
	tmap.Insert(us.world)
}

func (us *OppositionSystem) createPanes() {
	panelocal := viewers.NewPane(us.world)
	panelocal.AddBackground(engo.Point{X: 800, Y: 0})
	us.addDefaultOpposition(panelocal)
	us.pane = panelocal
	us.panes["LOCAL"] = panelocal

	panemap := viewers.NewPane(us.world)
	panemap.AddBackground(engo.Point{X: 800, Y: 0})
	us.panes["MAP"] = panemap
	panemap.Hide()
}

func (us *OppositionSystem) registerButton(t *ui.Text) {
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

func (us *OppositionSystem) addDefaultOpposition(pane *viewers.Pane) {
	text := ui.NewText("Nothing around.") // 240
	text.SetX(933).SetY(197)
	text.Insert(pane.World)
	pane.RegisterEntity(&text.BasicEntity, &text.RenderComponent)
}
