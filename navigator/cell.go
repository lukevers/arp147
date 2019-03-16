package navigator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lukevers/arp147/planet"
)

type Cell struct {
	X int64
	Y int64

	Planet *planet.Planet
}

func NewCell(cell *Cell) *Cell {
	if cell == nil {
		cell = &Cell{
			X: randInt(1234567, 7654321),
			Y: randInt(1234567, 7654321),
		}

	}

	cell.initialize()
	return cell
}

func (c *Cell) GetX() string {
	return fmt.Sprintf("0x%X", c.X)
}

func (c *Cell) GetY() string {
	return fmt.Sprintf("0x%X", c.Y)
}

func (c *Cell) HudX() string {
	return fmt.Sprintf("X: %s", c.GetX())
}

func (c *Cell) HudY() string {
	return fmt.Sprintf("Y: %s", c.GetY())
}

func (c *Cell) initialize() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	if randInt(0, 12)%12 == 0 {
		c.Planet = planet.New(planet.SizeViewer, planet.TypePlanet, true)
	}
}

func randInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}
