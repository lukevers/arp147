package text

type TextControlComponent struct {
	Mouse *Mouse
}

func (t *TextControlComponent) Type() string {
	return "TextControlComponent"
}
