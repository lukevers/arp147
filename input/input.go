package input

import (
	"strconv"

	"engo.io/engo"
)

var (
	sceneRegistry *[]Key
)

// Key contains the information needed for key events. It also contains three
// functions that can be set on each type of key event.
type Key struct {
	Name string
	Keys []engo.Key

	OnPress   func(engo.Key, *Modifiers)
	OnRelease func(engo.Key, *Modifiers)
}

// RegisterKeys registers keys with functions for the scene. It also handles
// registering modifiers to the global key management system so that we can
// tell if the key was used at the same time that a modifier key was used.
func RegisterKeys(keys []Key) {
	// Store the button registry globally
	sceneRegistry = &keys

	// Register all of the buttons
	for _, key := range keys {
		for index, k := range key.Keys {
			engo.Input.RegisterButton(
				key.Name+strconv.Itoa(index),
				k,
			)
		}
	}

	registerModifiers()
}
