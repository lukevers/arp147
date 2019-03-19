package navigator

import (
	"errors"

	"engo.io/engo"
)

type Map struct {
	Center *Cell
	Cells  map[int64]map[int64]*Cell
}

func NewMap() *Map {
	m := &Map{
		Cells: make(map[int64]map[int64]*Cell),
	}

	m.Center = m.GetCell(
		randInt(1234567, 7654321),
		randInt(1234567, 7654321),
	)

	m.Center.Discovered = true
	return m
}

func (m *Map) GoTo(x, y int64, force bool) error {
	if !force {
		if !m.InRange(x, y) {
			return errors.New("Could not jump: coordinates out of range.")
		}
	}

	m.Center = m.GetCell(x, y)
	m.Center.Discovered = true
	engo.Mailbox.Dispatch(MoveMessage{})
	return nil
}

func (m *Map) InRange(x, y int64) bool {
	for _, cell := range m.GetVisibleCells() {
		if cell.X == x && cell.Y == y {
			return true
		}
	}

	return false
}

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
