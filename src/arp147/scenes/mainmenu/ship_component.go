package mainmenu

type ShipComponent struct {
	Begin *bool
	Scale int
}

func (s *ShipComponent) Type() string {
	return "ShipComponent"
}
