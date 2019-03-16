package navigator

type Map struct {
	Center Cell
}

func NewMap() *Map {
	return &Map{
		Center: NewCell(),
	}
}

func (m *Map) GetVisibleCells() []Cell {
	return []Cell{
		Cell{
			X: m.Center.X - 1,
			Y: m.Center.Y - 1,
		},
		Cell{
			X: m.Center.X,
			Y: m.Center.Y - 1,
		},
		Cell{
			X: m.Center.X + 1,
			Y: m.Center.Y - 1,
		},
		Cell{
			X: m.Center.X - 1,
			Y: m.Center.Y,
		},
		m.Center,
		Cell{
			X: m.Center.X + 1,
			Y: m.Center.Y,
		},
		Cell{
			X: m.Center.X - 1,
			Y: m.Center.Y + 1,
		},
		Cell{
			X: m.Center.X,
			Y: m.Center.Y + 1,
		},
		Cell{
			X: m.Center.X + 1,
			Y: m.Center.Y + 1,
		},
	}
}
