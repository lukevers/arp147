package computer

import (
	"arp147/ui/fonts"
	"arp147/ui/position"
	"arp147/ui/text"
	"engo.io/engo"
	"image/color"
)

func (c *Computer) write(str string) {
	size := 16
	yoff := float32(c.line*size) * .9

	t := text.New(text.Text{
		Text:  str,
		Size:  float64(size),
		Font:  fonts.FONT_COMPUTER,
		Scale: engo.Point{1, 1},
		Color: text.Color{
			BG: color.Transparent,
			FG: color.White,
		},
		Position: position.Position{
			Point: engo.Point{
				X: float32(padding),
				Y: float32(padding) + yoff,
			},
			Position: position.TOP_LEFT,
		},
	})

	c.lines[c.line] = &line{
		text:   []*text.Text{t},
		locked: true,
	}

	c.world.AddEntity(t.Entity())

	// Increase the line number so we force a new line at the end of the output.
	c.line++
}
