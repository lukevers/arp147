package terminal

import (
	"github.com/lukevers/arp147/ui"
)

type line struct {
	text  []string
	chars []*ui.Text

	locked bool
}

func newLine() *line {
	l := &line{}

	// TOOD: autofill with prefix

	return l
}
