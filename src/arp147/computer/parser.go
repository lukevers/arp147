package computer

import (
	"strings"
)

func (c *Computer) parseLine(l *line) {
	// Generate a string of text from the line
	text := ""
	for _, char := range l.text {
		text += char.Text
	}

	// Remove duplicate spaces
	text = strings.Replace(text, "  ", " ", -1)

	// Split up line into arguments
	arguments := strings.Split(text, " ")

	// If for some reason we don't have any text at this point, stop.
	if len(arguments) < 1 {
		return
	}

	// Split out command and arguments
	command := arguments[0]
	if len(arguments) > 1 {
		arguments = arguments[1:]
	} else {
		arguments = []string{}
	}

	c.write(command)
}
