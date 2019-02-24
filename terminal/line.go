package terminal

import (
	"strings"

	"engo.io/engo"
	"github.com/lukevers/arp147/input"
	"github.com/lukevers/arp147/ui"
)

type line struct {
	text  []string
	chars []*ui.Text

	locked      bool
	prefixCount int
}

func (l *line) prefix(delegateKeyPress func(key engo.Key, mods *input.Modifiers)) {
	delegateKeyPress(engo.KeyA, &input.Modifiers{})
	delegateKeyPress(engo.KeyR, &input.Modifiers{})
	delegateKeyPress(engo.KeyP, &input.Modifiers{})
	delegateKeyPress(engo.KeyDash, &input.Modifiers{})
	delegateKeyPress(engo.KeyS, &input.Modifiers{})
	delegateKeyPress(engo.KeyH, &input.Modifiers{})
	delegateKeyPress(engo.KeyOne, &input.Modifiers{})
	delegateKeyPress(engo.KeyFour, &input.Modifiers{Shift: true})
	delegateKeyPress(engo.KeySpace, &input.Modifiers{})

	l.prefixCount = 9
}

func (l *line) evaluate(ts *TerminalSystem) {
	str := l.String()
	ts.command(str)
}

func (l *line) String() string {
	return strings.Join(
		l.text[l.prefixCount:],
		"",
	)
}
