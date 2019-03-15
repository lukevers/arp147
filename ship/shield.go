package ship

const (
	SheldLevelMin int = 0
)

type Shield struct {
	Level int
}

func (s *Shield) Increase(ship *Ship) {
	if s.Level < ship.defined.GetShieldLevelMax() {
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

func (s *Shield) Max(ship *Ship) {
	s.Level = ship.defined.GetShieldLevelMax()
	ship.SetSpriteCell(ship.defined.GetShieldLevelMax())
}

func (s *Shield) Min(ship *Ship) {
	s.Level = SheldLevelMin
	ship.SetSpriteCell(int(SheldLevelMin))
}
