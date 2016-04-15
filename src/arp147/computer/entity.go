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
		kc.On(engo.Key(i), c.handleKey)
	}

	// Numeric
	for i := engo.Zero; i <= engo.Nine; i++ {
		kc.On(engo.Key(i), c.handleKey)
	}

	// Misc
	kc.On(engo.Dash, c.handleKey)
	kc.On(engo.Apostrophe, c.handleKey)
	kc.On(engo.Semicolon, c.handleKey)
	kc.On(engo.Equals, c.handleKey)
	kc.On(engo.Comma, c.handleKey)
	kc.On(engo.Period, c.handleKey)
	kc.On(engo.Slash, c.handleKey)
	kc.On(engo.Backslash, c.handleKey)
	kc.On(engo.Backspace, c.handleKey)
	kc.On(engo.Tab, c.handleKey)
	//kc.On(engo.CapsLock, c.handleKey)
	kc.On(engo.Space, c.handleKey)
	kc.On(engo.Enter, c.handleKey)
	/*kc.On(engo.Escape, c.handleKey)
	kc.On(engo.ArrowLeft, c.handleKey)
	kc.On(engo.ArrowRight, c.handleKey)
	kc.On(engo.ArrowDown, c.handleKey)
	kc.On(engo.ArrowUp, c.handleKey)
	kc.On(engo.ArrowUp, c.handleKey)*/
	kc.On(engo.LeftBracket, c.handleKey)
	/*kc.On(engo.LeftShift, c.handleKey)
	kc.On(engo.LeftControl, c.handleKey)
	kc.On(engo.LeftSuper, c.handleKey)
	kc.On(engo.LeftAlt, c.handleKey)*/
	kc.On(engo.RightBracket, c.handleKey)
	/*kc.On(engo.RightShift, c.handleKey)
	kc.On(engo.RightControl, c.handleKey)
	kc.On(engo.RightSuper, c.handleKey)
	kc.On(engo.RightAlt, c.handleKey)*/

	// Add components
	c.entity.AddComponent(kc)
	c.entity.AddComponent(&engo.RenderComponent{})
	return c.entity
}

func (c *Computer) handleKey(key engo.Key) {
	size := 14
	var xoff, yoff float32

	// Catch special keys
	switch key {
	case engo.Enter:
		// An enter should advance us to the next line
		c.line++
		return
	case engo.Tab:
		// A tab should be translated into four spaces
		for i := 0; i < 4; i++ {
			c.handleKey(engo.Space)
		}
		return
	case engo.Backspace:
		// A backspace should delete the last character
		e := len(c.lines[c.line])
		if e > 0 {
			c.lines[c.line][e-1].Remove(c.world)
			c.lines[c.line] = c.lines[c.line][:e-1]
		}
		return
	}

	// Don't add any offset if we're on the very first character
	if len(c.lines) > 0 {
		// Don't add any x offset if we're the first character of the line
		if len(c.lines[c.line]) > 0 {
			xoff = float32(len(c.lines[c.line])*size) / 1.2
		}

		// Always create the y offset by the size of the font and the line
		yoff = float32(c.line * size)
	}

	// Create our character
	char := text.New(text.Text{
		Text:  string(key),
		Size:  float64(size),
		Font:  fonts.FONT_PRIMARY,
		Scale: engo.Point{1, 1},
		Color: text.Color{
			BG: color.Transparent,
			FG: color.White,
		},
		Position: position.Position{
			Point: engo.Point{
				X: 42 + xoff,
				Y: 42 + yoff,
			},
			Position: position.TOP_LEFT,
		},
	})

	// Add our character to the line
	c.lines[c.line] = append(c.lines[c.line], char)

	// Add it to the world
	c.world.AddEntity(char.Entity())
}
