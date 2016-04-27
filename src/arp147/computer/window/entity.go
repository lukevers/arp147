package window

import (
	"arp147/computer/line"
	"arp147/systems/key"
	"arp147/ui/fonts"
	"arp147/ui/position"
	"arp147/ui/text"
	"engo.io/ecs"
	"engo.io/engo"
	"image/color"
)

func (w *Window) Entity() *ecs.Entity {
	w.entity = ecs.NewEntity("RenderSystem", "KeySystem")

	kc := &key.KeyComponent{}

	// Alpha
	for i := engo.A; i <= engo.Z; i++ {
		kc.On(engo.Key(i), w.printKey)
	}

	// Numeric
	for i := engo.Zero; i <= engo.Nine; i++ {
		kc.On(engo.Key(i), w.printKey)
	}

	// Misc
	kc.On(engo.Dash, w.printKey)
	kc.On(engo.Apostrophe, w.printKey)
	kc.On(engo.Semicolon, w.printKey)
	kc.On(engo.Equals, w.printKey)
	kc.On(engo.Comma, w.printKey)
	kc.On(engo.Period, w.printKey)
	kc.On(engo.Slash, w.printKey)
	kc.On(engo.Backslash, w.printKey)
	kc.On(engo.Backspace, w.printKey)
	kc.On(engo.Tab, w.printKey)
	//kc.On(engo.CapsLock, w.printKey)
	kc.On(engo.Space, w.printKey)
	kc.On(engo.Enter, w.printKey)
	kc.On(engo.Escape, w.printKey)
	/*kc.On(engo.ArrowLeft, w.printKey)
	kc.On(engo.ArrowRight, w.printKey)
	kc.On(engo.ArrowDown, w.printKey)
	kc.On(engo.ArrowUp, w.printKey)
	kc.On(engo.ArrowUp, w.printKey)*/
	kc.On(engo.LeftBracket, w.printKey)
	/*kc.On(engo.LeftShift, w.printKey)
	kc.On(engo.LeftControl, w.printKey)
	kc.On(engo.LeftSuper, w.printKey)
	kc.On(engo.LeftAlt, w.printKey)*/
	kc.On(engo.RightBracket, w.printKey)
	/*kc.On(engo.RightShift, w.printKey)
	kc.On(engo.RightControl, w.printKey)
	kc.On(engo.RightSuper, w.printKey)
	kc.On(engo.RightAlt, w.printKey)*/

	// Add components
	w.entity.AddComponent(kc)
	w.entity.AddComponent(&engo.RenderComponent{})
	return w.entity
}

func (w *Window) printKey(key engo.Key, caps bool) {
	// If the computer is not active, don't continue with anything.
	if !w.Active {
		return
	}

	size := 16
	var xoff, yoff float32

	// Initialize struct if it doesn't exist yet
	if w.Lines[w.Line] == nil {
		w.Lines[w.Line] = &line.Line{}
	}

	// Catch special keys
	switch key {
	case engo.Escape:
		w.StopSession()
		return
	case engo.Enter:
		// An enter should advance us to the next line
		w.Lines[w.Line].Locked = true
		w.Line++

		// Make sure that there was at least one character printed before
		// parsing the line.
		if len(w.Lines[w.Line-1].Text) > 0 {
			// Send the line to be parsed into commands and arguments and then
			// dispatch the command and arguments to be run.
			cargs := w.Lines[w.Line-1].ToCArgs()
			if cargs != nil {
				err := w.Terminal.Dispatch(cargs)
				if err != nil {
					w.Write(err.Error())
				}
			}
		}
		return
	case engo.Tab:
		// A tab should be translated into four spaces
		for i := 0; i < 4; i++ {
			w.printKey(engo.Space, caps)
		}
		return
	case engo.Backspace:
		// A backspace should delete the last character
		e := len(w.Lines[w.Line].Text)
		if e > 0 {
			w.Lines[w.Line].Text[e-1].Remove(w.world)
			w.Lines[w.Line].Text = w.Lines[w.Line].Text[:e-1]
		} else {
			if w.Line > 0 && !w.Lines[w.Line-1].Locked {
				w.Line--
				e = len(w.Lines[w.Line].Text)
				w.Lines[w.Line].Text[e-1].Remove(w.world)
				w.Lines[w.Line].Text = w.Lines[w.Line].Text[:e-1]
			}
		}

		return
	}

	// Don't add any offset if we're on the very first character
	if len(w.Lines) > 0 {
		// Don't add any x offset if we're the first character of the line
		if len(w.Lines[w.Line].Text) > 0 {
			xoff = float32(len(w.Lines[w.Line].Text)*size) * .5
			if xoff >= (engo.Width() - float32(padding*2)) {
				xoff = 0
				w.Lines[w.Line].Locked = false
				w.Line++
			}
		}

		// Always create the y offset by the size of the font and the line
		yoff = float32(w.Line*size) * .9
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
	if w.Lines[w.Line] == nil {
		w.Lines[w.Line] = &line.Line{}
	}

	w.Lines[w.Line].Text = append(w.Lines[w.Line].Text, char)

	// Add it to the world
	w.world.AddEntity(char.Entity())
}
