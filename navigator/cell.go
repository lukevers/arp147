package navigator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lukevers/arp147/entity"
)

// Cell defines an X,Y location, and also contains data about what is in that
// location.
type Cell struct {
	X int64
	Y int64

	Entity     *entity.Entity
	Discovered bool
}

// NewCell initializes cell data from a given *Cell.
func NewCell(cell *Cell) *Cell {
	cell.initialize()
	return cell
}

// GetX returns a hex formatted X location
func (c *Cell) GetX() string {
	return fmt.Sprintf("0x%X", c.X)
}

// GetY returns a hex formatted Y location
func (c *Cell) GetY() string {
	return fmt.Sprintf("0x%X", c.Y)
}

// HudX returns an HUD friendly X location
func (c *Cell) HudX() string {
	return fmt.Sprintf("X: %s", c.GetX())
}

// HudY returns an HUD friendly Y location
func (c *Cell) HudY() string {
	return fmt.Sprintf("Y: %s", c.GetY())
}

func (c *Cell) initialize() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	if true {
		// if chance(6) {
		c.Entity = entity.New(entity.SizeViewer, entity.TypePlanet, true)
	}
}

func randInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func chance(r int64) bool {
	return randInt(0, r)%12 == 0
}
