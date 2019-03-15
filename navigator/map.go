package navigator

import (
	"fmt"
	"math/rand"
)

type Map struct {
	X int64
	Y int64
}

func NewMap() *Map {
	return &Map{
		X: randInt(1234567, 7654321),
		Y: randInt(1234567, 7654321),
	}
}

func (m *Map) GetX() string {
	return fmt.Sprintf("0x%X", m.X)
}

func (m *Map) GetY() string {
	return fmt.Sprintf("0x%X", m.Y)
}

func (m *Map) HudX() string {
	return fmt.Sprintf("X: %s", m.GetX())
}

func (m *Map) HudY() string {
	return fmt.Sprintf("Y: %s", m.GetY())
}

func randInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}
