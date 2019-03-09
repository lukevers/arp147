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
func (os *OppositionSystem) New(w *ecs.World) {
	os.world = w
	os.panes = make(map[string]*viewers.Pane)

	os.addButtons()
	os.createPanes()

	log.Println("OppositionSystem initialized")
}

func (os *OppositionSystem) SetActiveTab(tab *ui.Text) {
	os.tab = tab
}

func (os *OppositionSystem) GetActiveTab() *ui.Text {
	return os.tab
}

func (os *OppositionSystem) GetTabs() []*ui.Text {
	return os.tabs
}

func (os *OppositionSystem) GetActivePane() *viewers.Pane {
	return os.pane
}

func (os *OppositionSystem) SetActivePane(pane *viewers.Pane) {
	os.pane = pane
}

func (os *OppositionSystem) GetPanes() map[string]*viewers.Pane {
	return os.panes
}

func (os *OppositionSystem) addButtons() {
	tlocal := ui.NewText("LOCAL")
	tlocal.Font.Size = 12
	tlocal.SetX(820).SetY(16)
	os.tabs = append(os.tabs, tlocal)
	os.tab = tlocal
	// os.registerButton(tlocal)
	viewers.RegisterButton(tlocal, os)
	tlocal.Insert(os.world)

	tmap := ui.NewText("MAP")
	tmap.Font.Size = 12
	tmap.Font.FG = color.Alpha16{0x666F}
	tmap.SetX(870).SetY(16)
	os.tabs = append(os.tabs, tmap)
	viewers.RegisterButton(tmap, os)
	// os.registerButton(tmap)
	tmap.Insert(os.world)
}

func (os *OppositionSystem) createPanes() {
	panelocal := viewers.NewPane(os.world)
	panelocal.AddBackground(engo.Point{X: 800, Y: 0})
	os.addDefaultOpposition(panelocal)
	os.pane = panelocal
	os.panes["LOCAL"] = panelocal

	panemap := viewers.NewPane(os.world)
	panemap.AddBackground(engo.Point{X: 800, Y: 0})
	os.panes["MAP"] = panemap
	panemap.Hide()
}

// func (os *OppositionSystem) registerButton(t *ui.Text) {
// 	t.OnClicked(func(entity *ecs.BasicEntity, dt float32) {
// 		if os.tab == t {
// 			return
// 		}

// 		for _, tab := range os.tabs {
// 			tab.Font.FG = color.Alpha16{0x666F}
// 		}

// 		for _, pane := range os.panes {
// 			pane.Hide()
// 		}

// 		t.Font.FG = color.White

// 		os.tab = t
// 		os.pane = os.panes[t.Text]
// 		os.pane.Show()
// 	}).OnEnter(func(entity *ecs.BasicEntity, dt float32) {
// 		if os.tab == t {
// 			return
// 		}

// 		t.Font.FG = color.Alpha16{0xAAAF}
// 	}).OnLeave(func(entity *ecs.BasicEntity, dt float32) {
// 		if os.tab == t {
// 			return
// 		}

// 		t.Font.FG = color.Alpha16{0x666F}
// 	})
// }

func (os *OppositionSystem) addDefaultOpposition(pane *viewers.Pane) {
	text := ui.NewText("Nothing around.") // 240
	text.SetX(933).SetY(197)
	text.Insert(pane.World)
	pane.RegisterEntity(&text.BasicEntity, &text.RenderComponent)
}
