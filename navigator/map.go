package navigator

import (
	"errors"
	"math/rand"
	"time"

	"engo.io/engo"
)

// Map defines a world map.
type Map struct {
	Center *Cell
	Cells  map[int64]map[int64]*Cell
}

// NewMap generates a new world map.
func NewMap() *Map {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	m := &Map{
		Cells: make(map[int64]map[int64]*Cell),
	}

	m.Center = m.GetCell(
		randInt(123, 765),
		randInt(123, 765),
	)

	m.Center.Discovered = true
	return m
}

// GoTo allows a player to jump if they can, or if force is true.
func (m *Map) GoTo(x, y int64, force bool) error {
	if !force {
		if !m.InRange(x, y) {
			return errors.New("could not jump: coordinates out of range")
		}
	}

	m.Center = m.GetCell(x, y)
	m.Center.Discovered = true
	engo.Mailbox.Dispatch(MoveMessage{})
	return nil
}

// InRange checks if coordinates are in range to jump or not.
func (m *Map) InRange(x, y int64) bool {
	for _, cell := range m.GetVisibleCells() {
		if cell.X == x && cell.Y == y {
			return true
		}
	}

	return false
}

// GetCell looks for a Cell, and if it does not yet exist in the world map,
// it generates it.
func (m *Map) GetCell(x, y int64) *Cell {
	if _, exists := m.Cells[x]; !exists {
		m.Cells[x] = make(map[int64]*Cell)
	}

	if cell, exists := m.Cells[x][y]; exists {
		return cell
	}

	m.Cells[x][y] = NewCell(&Cell{X: x, Y: y})
	return m.Cells[x][y]
}

// GetVisibleCells returns a slice of jumpable locations in range.
func (m *Map) GetVisibleCells() []*Cell {
	return []*Cell{
		m.GetCell(m.Center.X-1, m.Center.Y-1),
		m.GetCell(m.Center.X, m.Center.Y-1),
		m.GetCell(m.Center.X+1, m.Center.Y-1),
		m.GetCell(m.Center.X-1, m.Center.Y),
		m.Center,
		m.GetCell(m.Center.X+1, m.Center.Y),
		m.GetCell(m.Center.X-1, m.Center.Y+1),
		m.GetCell(m.Center.X, m.Center.Y+1),
		m.GetCell(m.Center.X+1, m.Center.Y+1),
	}
}
