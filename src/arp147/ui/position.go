package ui

import (
	"engo.io/engo"
)

const (
	PositionTopLeft = 1 << iota
	PositionTopCenter
	PositionTopRight
	PositionCenterLeft
	PositionCenterCenter
	PositionCenterRight
	PositionBottomLeft
	PositionBottomCenter
	PositionBottomRight
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
	width := engo.GameWidth()
	height := engo.GameHeight()

	switch p.Position {
	case PositionTopLeft:
	case PositionTopCenter:
		p.Point.X = ((width / 2) - p.Point.X) - (dx / 2)
	case PositionTopRight:
		p.Point.X = (width - p.Point.X) - dx
	case PositionCenterLeft:
		p.Point.Y = (height / 2) - (dy / 2)
	case PositionCenterCenter:
		p.Point.X = (width / 2) - (dx / 2)
		p.Point.Y = (height / 2) - (dy / 2)
	case PositionCenterRight:
		p.Point.Y = (height / 2) - (dy / 2)
		p.Point.X = (width - p.Point.X) - dx
	case PositionBottomLeft:
		p.Point.Y = (height - p.Point.Y) - dy
	case PositionBottomCenter:
		p.Point.X = (width / 2) - (dx / 2)
		p.Point.Y = (height - p.Point.Y) - dy
	case PositionBottomRight:
		p.Point.X = (width - p.Point.X) - dx
		p.Point.Y = (height - p.Point.Y) - dy
	}

	return p.Point
}
