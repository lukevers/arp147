package display

type FakePlayerComponent struct {
	X, Y  float32
	Begin *bool
}

func (f *FakePlayerComponent) Type() string {
	return "FakePlayerComponent"
}
