package ship

const (
	SheldLevelMin SheldLevel = 0
	SheldLevelMax SheldLevel = 3
)

type (
	SheldLevel int
)

type Shield struct {
	Level SheldLevel
}

func (s *Shield) Increase(ship *Ship) {
	if s.Level < SheldLevelMax {
		s.Level++
		ship.SetSpriteCell(int(s.Level))
	}
}

func (s *Shield) Decrease(ship *Ship) {
	if s.Level > SheldLevelMin {
		s.Level--
		ship.SetSpriteCell(int(s.Level))
	}
}
