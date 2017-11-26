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
		engo.LeftAlt,
		engo.RightAlt,
	)

	// Register the control keys
	engo.Input.RegisterButton(
		ModifierControl,
		engo.LeftControl,
		engo.RightControl,
	)

	// Register the shift keys
	engo.Input.RegisterButton(
		ModifierShift,
		engo.LeftShift,
		engo.RightShift,
	)

	// Register the super keys
	engo.Input.RegisterButton(
		ModifierSuper,
		engo.LeftSuper,
		engo.RightSuper,
	)
}
