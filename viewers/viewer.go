package viewers

import (
	"github.com/lukevers/arp147/ui"
)

type Viewer interface {
	SetActiveTab(*ui.Text)
	GetActiveTab() *ui.Text
	GetTabs() []*ui.Text

	GetActivePane() *Pane
	SetActivePane(*Pane)
	GetPanes() map[string]*Pane
}
