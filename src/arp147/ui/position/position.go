package position

import (
	"github.com/engoengine/engo"
)

const (
	TOP_LEFT = 1 << iota
	TOP_CENTER
	TOP_RIGHT
	CENTER_LEFT
	CENTER_CENTER
	CENTER_RIGHT
	BOTTOM_LEFT
	BOTTOM_CENTER
	BOTTOM_RIGHT
)

type Position struct {
	Point    engo.Point
	Position int
}

// Calculate takes the position and converts it to an engo.Point struct that
// also is aware of the location on the screen the entity is supposed to
// display at. The parameters in this function are the width (dx) and the
// height (dy) of the entity. Also, p.Point contains a point which is default
// set for TOP_LEFT with margin to where the point should be located. If the
// entity has been set with specific margin (aka the default point), that
// should also be taken into consideration for figuring out the location of the
// point; however, if a point is set for CENTER_CENTER, the margin should be
// discarded to avoid non-CENTER_CENTER behavior. This also runs true for other
// points where CENTER is involved.
func (p Position) Calculate(dx, dy float32) engo.Point {
	width := engo.WindowWidth()
	height := engo.WindowHeight()

	switch p.Position {
	case TOP_LEFT:
	case TOP_CENTER:
		p.Point.X = ((width / 2) - p.Point.X) - (dx / 2)
	case TOP_RIGHT:
		p.Point.X = (width - p.Point.X) - dx
	case CENTER_LEFT:
		p.Point.Y = (height / 2) - (dy / 2)
	case CENTER_CENTER:
		p.Point.X = (width / 2) - (dx / 2)
		p.Point.Y = (height / 2) - (dy / 2)
	case CENTER_RIGHT:
		p.Point.Y = (height / 2) - (dy / 2)
		p.Point.X = (width - p.Point.X) - dx
	case BOTTOM_LEFT:
		p.Point.Y = (height - p.Point.Y) - dy
	case BOTTOM_CENTER:
		p.Point.X = (width / 2) - (dx / 2)
		p.Point.Y = (height - p.Point.Y) - dy
	case BOTTOM_RIGHT:
		p.Point.X = (width - p.Point.X) - dx
		p.Point.Y = (height - p.Point.Y) - dy
	}

	return p.Point
}
