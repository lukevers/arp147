package input

import (
	"engo.io/engo"
)

var registry *[]Key

// Key contains the information needed for key events. It also contains three
// functions that can be set on each type of key event.
type Key struct {
	Name string
	Keys []engo.Key

	Down         func()
	JustPressed  func()
	JustReleased func()
}

// RegisterButtons registers buttons with functions for the scene.
func RegisterButtons(buttons []Key) {
	// Store the button registry globally
	registry = &buttons

	// Register all of the buttons
	for _, key := range buttons {
		engo.Input.RegisterButton(
			key.Name,
			key.Keys...,
		)
	}
}
