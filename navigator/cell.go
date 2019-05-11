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

	Planet     *planet.Planet
	Discovered bool
}

func NewCell(cell *Cell) *Cell {
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

	if chance(6) {
		c.Planet = planet.New(planet.SizeViewer, planet.TypePlanet, true)
	}
}

func randInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func chance(r int64) bool {
	return randInt(0, r)%12 == 0
}
