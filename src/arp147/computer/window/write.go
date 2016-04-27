package window

import (
	"arp147/computer/line"
	"arp147/ui/fonts"
	"arp147/ui/position"
	"arp147/ui/text"
	"engo.io/engo"
	"image/color"
)

// Write a new line to the window
func (w *Window) NewLine() {
	w.Lines[w.Line] = &line.Line{
		Text:   []*text.Text{},
		Locked: true,
	}

	w.Line++
}

// Write a block of text to the window
func (w *Window) Write(str string) {
	size := 16
	yoff := float32(w.Line*size) * .9

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

	w.Lines[w.Line] = &line.Line{
		Text:   []*text.Text{t},
		Locked: true,
	}

	w.world.AddEntity(t.Entity())

	// Increase the line number so we force a new line at the end of the output.
	w.Line++
}
