package input

import (
	"engo.io/engo"
)

var (
	hasRegisteredModifiers bool
)

const (
	ModifierAlt     = "alt"
	ModifierControl = "control"
	ModifierShift   = "shift"
	ModifierSuper   = "super"
)

// Modifiers contains a bool for every kind of modifier key that could be used
// at the same time as the key being used.
type Modifiers struct {
	Alt     bool
	Control bool
	Shift   bool
	Super   bool
}

func registerModifiers() {
	// If we've already registered the modifiers, we don't need to do it again.
	if hasRegisteredModifiers {
		return
	}

	// Make sure we don't re-register later.
	hasRegisteredModifiers = true

	// Register the alt keys
	engo.Input.RegisterButton(
		ModifierAlt,
		engo.KeyLeftAlt,
		engo.KeyRightAlt,
	)

	// Register the control keys
	engo.Input.RegisterButton(
		ModifierControl,
		engo.KeyLeftControl,
		engo.KeyRightControl,
	)

	// Register the shift keys
	engo.Input.RegisterButton(
		ModifierShift,
		engo.KeyLeftShift,
		engo.KeyRightShift,
	)

	// Register the super keys
	engo.Input.RegisterButton(
		ModifierSuper,
		engo.KeyLeftSuper,
		engo.KeyRightSuper,
	)
}
