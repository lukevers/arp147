package computer

import (
	"arp147/systems/key"
	"arp147/ui/fonts"
	"arp147/ui/position"
	"arp147/ui/text"
	"engo.io/ecs"
	"engo.io/engo"
	"image/color"
)

func (c *Computer) Entity() *ecs.Entity {
	c.entity = ecs.NewEntity("RenderSystem", "KeySystem")

	kc := &key.KeyComponent{}

	// Alpha
	for i := engo.A; i <= engo.Z; i++ {
		kc.On(engo.Key(i), c.printKey)
	}

	// Numeric
	for i := engo.Zero; i <= engo.Nine; i++ {
		kc.On(engo.Key(i), c.printKey)
	}

	// Misc
	kc.On(engo.Dash, c.printKey)
	kc.On(engo.Apostrophe, c.printKey)
	kc.On(engo.Semicolon, c.printKey)
	kc.On(engo.Equals, c.printKey)
	kc.On(engo.Comma, c.printKey)
	kc.On(engo.Period, c.printKey)
	kc.On(engo.Slash, c.printKey)
	kc.On(engo.Backslash, c.printKey)
	kc.On(engo.Backspace, c.printKey)
	kc.On(engo.Tab, c.printKey)
	//kc.On(engo.CapsLock, c.printKey)
	kc.On(engo.Space, c.printKey)
	kc.On(engo.Enter, c.printKey)
	kc.On(engo.Escape, c.printKey)
	/*kc.On(engo.ArrowLeft, c.printKey)
	kc.On(engo.ArrowRight, c.printKey)
	kc.On(engo.ArrowDown, c.printKey)
	kc.On(engo.ArrowUp, c.printKey)
	kc.On(engo.ArrowUp, c.printKey)*/
	kc.On(engo.LeftBracket, c.printKey)
	/*kc.On(engo.LeftShift, c.printKey)
	kc.On(engo.LeftControl, c.printKey)
	kc.On(engo.LeftSuper, c.printKey)
	kc.On(engo.LeftAlt, c.printKey)*/
	kc.On(engo.RightBracket, c.printKey)
	/*kc.On(engo.RightShift, c.printKey)
	kc.On(engo.RightControl, c.printKey)
	kc.On(engo.RightSuper, c.printKey)
	kc.On(engo.RightAlt, c.printKey)*/

	// Add components
	c.entity.AddComponent(kc)
	c.entity.AddComponent(&engo.RenderComponent{})
	return c.entity
}

func (c *Computer) printKey(key engo.Key, caps bool) {
	// If the computer is not active, don't continue with anything.
	if !c.Active {
		return
	}

	size := 16
	var xoff, yoff float32

	// Initialize struct if it doesn't exist yet
	if c.lines[c.line] == nil {
		c.lines[c.line] = &line{}
	}

	// Catch special keys
	switch key {
	case engo.Escape:
		c.StopSession()
		return
	case engo.Enter:
		// An enter should advance us to the next line
		c.lines[c.line].locked = true
		c.line++

		// Make sure that there was at least one character printed before
		// parsing the line.
		if len(c.lines[c.line-1].text) > 0 {
			// Send the line to be parsed into commands and arguments and then
			// dispatch the command and arguments to be run.
			c.parseLine(c.lines[c.line-1])
		}
		return
	case engo.Tab:
		// A tab should be translated into four spaces
		for i := 0; i < 4; i++ {
			c.printKey(engo.Space, caps)
		}
		return
	case engo.Backspace:
		// A backspace should delete the last character
		e := len(c.lines[c.line].text)
		if e > 0 {
			c.lines[c.line].text[e-1].Remove(c.world)
			c.lines[c.line].text = c.lines[c.line].text[:e-1]
		} else {
			if c.line > 0 && !c.lines[c.line-1].locked {
				c.line--
				e = len(c.lines[c.line].text)
				c.lines[c.line].text[e-1].Remove(c.world)
				c.lines[c.line].text = c.lines[c.line].text[:e-1]
			}
		}

		return
	}

	// Don't add any offset if we're on the very first character
	if len(c.lines) > 0 {
		// Don't add any x offset if we're the first character of the line
		if len(c.lines[c.line].text) > 0 {
			xoff = float32(len(c.lines[c.line].text)*size) * .5
			if xoff >= (engo.Width() - float32(padding*2)) {
				xoff = 0
				c.lines[c.line].locked = false
				c.line++
			}
		}

		// Always create the y offset by the size of the font and the line
		yoff = float32(c.line*size) * .9
	}

	if !caps {
		// If the key we've pushed is [a-z], make it lowercase.
		if key >= engo.A && key <= engo.Z {
			key += 32
		}
	} else {
		switch key {
		// Numbers
		case engo.One:
			key = 33
		case engo.Two:
			key = 64
		case engo.Three:
			key = 35
		case engo.Four:
			key = 36
		case engo.Five:
			key = 37
		case engo.Six:
			key = 94
		case engo.Seven:
			key = 38
		case engo.Eight:
			key = 42
		case engo.Nine:
			key = 40
		case engo.Zero:
			key = 41

		// Misc
		case engo.Dash:
			key = 95
		case engo.Equals:
			key = 43
		case engo.LeftBracket:
			key = 123
		case engo.RightBracket:
			key = 125
		case engo.Backslash:
			key = 124
		case engo.Semicolon:
			key = 58
		case engo.Apostrophe:
			key = 34
		case engo.Comma:
			key = 60
		case engo.Period:
			key = 62
		case engo.Slash:
			key = 63
		}

		if key >= engo.Zero && key <= engo.Nine {
			key -= 16
		}
	}

	// Create our character
	char := text.New(text.Text{
		Text:  string(key),
		Size:  float64(size),
		Font:  fonts.FONT_COMPUTER,
		Scale: engo.Point{1, 1},
		Color: text.Color{
			BG: color.Transparent,
			FG: color.White,
		},
		Position: position.Position{
			Point: engo.Point{
				X: float32(padding) + xoff,
				Y: float32(padding) + yoff,
			},
			Position: position.TOP_LEFT,
		},
	})

	// Add our character to the line
	if c.lines[c.line] == nil {
		c.lines[c.line] = &line{}
	}

	c.lines[c.line].text = append(c.lines[c.line].text, char)

	// Add it to the world
	c.world.AddEntity(char.Entity())
}
