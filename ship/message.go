package ship

type NewShipMessage struct {
	Ship *Ship
}

func (NewShipMessage) Type() string {
	return "NewShipMessage"
}
