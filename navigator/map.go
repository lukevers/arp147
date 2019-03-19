package navigator

import (
	"engo.io/engo"
)

type Map struct {
	Center *Cell
	Cells  map[int64]map[int64]*Cell
}

func NewMap() *Map {
	m := &Map{
		Center: NewCell(nil),
		Cells:  make(map[int64]map[int64]*Cell),
	}

	m.Center.Discovered = true
	return m
}

func (m *Map) GoTo(x, y int64) {
	m.Center = m.GetCell(x, y)
	m.Center.Discovered = true
	engo.Mailbox.Dispatch(MoveMessage{})
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
