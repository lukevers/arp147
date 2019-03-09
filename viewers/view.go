package viewers

import (
	"image/color"

	"engo.io/ecs"
	"github.com/lukevers/arp147/ui"
)

type View struct {
	tab  *ui.Text
	tabs []*ui.Text

	pane  *Pane
	panes map[string]*Pane
}

func New() *View {
	return &View{
		panes: make(map[string]*Pane),
	}
}

func (v *View) SetActiveTab(tab *ui.Text) {
	v.tab = tab
}

func (v *View) GetActiveTab() *ui.Text {
	return v.tab
}

func (v *View) GetTabs() []*ui.Text {
	return v.tabs
}

func (v *View) GetActivePane() *Pane {
	return v.pane
}

func (v *View) SetActivePane(pane *Pane) {
	v.pane = pane
}

func (v *View) GetPanes() map[string]*Pane {
	return v.panes
}

func (v *View) AddPane(name string, pane *Pane) {
	v.panes[name] = pane
}

func (v *View) AddTab(tab *ui.Text) {
	v.tabs = append(v.tabs, tab)
}

func (v *View) RegisterButton(button *ui.Text) {
	v.AddTab(button)

	button.OnClicked(func(entity *ecs.BasicEntity, dt float32) {
		panes := v.GetPanes()
		if v.GetActiveTab() == button {
			return
		}

		for _, tab := range v.GetTabs() {
			tab.Font.FG = color.Alpha16{0x666F}
		}

		for _, pane := range panes {
			pane.Hide()
		}

		button.Font.FG = color.White

		v.SetActiveTab(button)
		v.SetActivePane(panes[button.Text])
		v.GetActivePane().Show()
	}).OnEnter(func(entity *ecs.BasicEntity, dt float32) {
		if v.GetActiveTab() == button {
			return
		}

		button.Font.FG = color.Alpha16{0xAAAF}
	}).OnLeave(func(entity *ecs.BasicEntity, dt float32) {
		if v.GetActiveTab() == button {
			return
		}

		button.Font.FG = color.Alpha16{0x666F}
	})
}
