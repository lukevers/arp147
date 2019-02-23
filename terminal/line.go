package terminal

import (
	"github.com/lukevers/arp147/ui"
)

type line struct {
	text  []string
	chars []*ui.Text

	locked bool
}
