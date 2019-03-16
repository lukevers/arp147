package navigator

import (
	"fmt"
	"math/rand"
)

type Cell struct {
	X int64
	Y int64
}

func NewCell() Cell {
	return Cell{
		X: randInt(1234567, 7654321),
		Y: randInt(1234567, 7654321),
	}
}

func (c Cell) GetX() string {
	return fmt.Sprintf("0x%X", c.X)
}

func (c Cell) GetY() string {
	return fmt.Sprintf("0x%X", c.Y)
}

func (c Cell) HudX() string {
	return fmt.Sprintf("X: %s", c.GetX())
}

func (c Cell) HudY() string {
	return fmt.Sprintf("Y: %s", c.GetY())
}

func randInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}
