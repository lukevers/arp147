package mainmenu

type ShipComponent struct {
	Begin *bool
	speed float32
}

func (s *ShipComponent) Type() string {
	return "ShipComponent"
}
