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
	delegateKeyPress(input.StringToKey("a"))
	delegateKeyPress(input.StringToKey("r"))
	delegateKeyPress(input.StringToKey("p"))
	delegateKeyPress(input.StringToKey("-"))
	delegateKeyPress(input.StringToKey("s"))
	delegateKeyPress(input.StringToKey("h"))
	delegateKeyPress(input.StringToKey("1"))
	delegateKeyPress(input.StringToKey("$"))
	delegateKeyPress(input.StringToKey(" "))

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
