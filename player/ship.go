package player

type Ship struct {
	Terminal *Terminal
}

func NewShip() *Ship {
	return &Ship{
		Terminal: NewTerminal(),
	}
}
