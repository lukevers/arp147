package navigator

// MoveMessage is the message that the navigator package emits when a player
// moves.
type MoveMessage struct{}

// Type defines the stringified name of the message.
func (MoveMessage) Type() string {
	return "MoveMessage"
}
