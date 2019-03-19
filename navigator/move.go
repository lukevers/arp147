package navigator

type MoveMessage struct{}

func (MoveMessage) Type() string {
	return "MoveMessage"
}
