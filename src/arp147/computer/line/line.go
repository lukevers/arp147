package line

import (
	"arp147/ui/text"
	"strings"
)

type Line struct {
	Text   []*text.Text
	Locked bool
}

// Convert a line into a command and arguments
func (l *Line) ToCArgs() *Command {
	// Generate a string of text from the line
	text := ""
	for _, char := range l.Text {
		text += char.Text
	}

	// Remove duplicate spaces
	text = strings.Replace(text, "  ", " ", -1)

	// Split up line into arguments
	arguments := strings.Split(text, " ")

	// If for some reason we don't have any text at this point, stop.
	if len(arguments) < 1 {
		return nil
	}

	// Split out command and arguments
	command := arguments[0]
	if len(arguments) > 1 {
		arguments = arguments[1:]
	} else {
		arguments = []string{}
	}

	return &Command{
		Command:   command,
		Arguments: arguments,
	}
}
