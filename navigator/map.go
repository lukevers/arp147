package navigator

type Map struct {
	Center *Cell
}

func NewMap() *Map {
	return &Map{
		Center: NewCell(nil),
	}
}

func (m *Map) GetVisibleCells() []*Cell {
	return []*Cell{
		NewCell(&Cell{
			X: m.Center.X - 1,
			Y: m.Center.Y - 1,
		}),
		NewCell(&Cell{
			X: m.Center.X,
			Y: m.Center.Y - 1,
		}),
		NewCell(&Cell{
			X: m.Center.X + 1,
			Y: m.Center.Y - 1,
		}),
		NewCell(&Cell{
			X: m.Center.X - 1,
			Y: m.Center.Y,
		}),
		m.Center,
		NewCell(&Cell{
			X: m.Center.X + 1,
			Y: m.Center.Y,
		}),
		NewCell(&Cell{
			X: m.Center.X - 1,
			Y: m.Center.Y + 1,
		}),
		NewCell(&Cell{
			X: m.Center.X,
			Y: m.Center.Y + 1,
		}),
		NewCell(&Cell{
			X: m.Center.X + 1,
			Y: m.Center.Y + 1,
		}),
	}
}
