package input

import (
	"engo.io/engo"
)

var (
	hasRegisteredModifiers bool
)

// Defines modifiers that can be used.
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

	Ignore bool
	Output bool
	Redraw bool
	Line   *string
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

// StringToKey takes a string and converts it to an engo.Key.
func StringToKey(str string, setmods ...*Modifiers) (engo.Key, *Modifiers) {
	// Defaults
	key := engo.KeySpace
	mods := &Modifiers{}

	if len(setmods) > 0 {
		mods = setmods[0]
	}

	bytes := []byte(str)
	if len(bytes) < 1 {
		return key, mods
	}

	key = engo.Key(bytes[0])
	return key, mods
}

// KeyToString converts an engo.Key to a string.
func KeyToString(key engo.Key, mods *Modifiers) (symbol string) {
	// // If the key is [a-z] apply shift rules.
	if key >= engo.KeyA && key <= engo.KeyZ {
		if mods.Shift {
			symbol = string(key)
		} else {
			symbol = string(key + 32)
		}
	} else {
		// Convert non [a-z] letters when shift is used
		if mods.Shift {
			switch key {
			case engo.KeyZero:
				symbol = ")"
			case engo.KeyOne:
				symbol = "!"
			case engo.KeyTwo:
				symbol = "@"
			case engo.KeyThree:
				symbol = "#"
			case engo.KeyFour:
				symbol = "$"
			case engo.KeyFive:
				symbol = "%"
			case engo.KeySix:
				symbol = "^"
			case engo.KeySeven:
				symbol = "&"
			case engo.KeyEight:
				symbol = "*"
			case engo.KeyNine:
				symbol = "("

			case engo.KeyDash:
				symbol = "_"
			case engo.KeyGrave:
				symbol = "~"
			case engo.KeyApostrophe:
				symbol = "\""
			case engo.KeySemicolon:
				symbol = ":"
			case engo.KeyEquals:
				symbol = "+"
			case engo.KeyComma:
				symbol = "<"
			case engo.KeyPeriod:
				symbol = ">"
			case engo.KeySlash:
				symbol = "?"
			case engo.KeyBackslash:
				symbol = "|"
			case engo.KeyLeftBracket:
				symbol = "{"
			case engo.KeyRightBracket:
				symbol = "}"
			default:
				symbol = string(key)
			}

			// TODO
			//   - see above
			//   - will this be different for different keyboard layouts?
			//     - we can't assume everyone uses US QWERTY
			//	   - I should learn how other layouts work
		} else {
			// Otherwise we just use the actual key here
			symbol = string(key)
		}
	}

	return
}
