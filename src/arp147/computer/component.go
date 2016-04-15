package computer

type ComputerComponent struct {
	Visible bool
}

func (c *ComputerComponent) Type() string {
	return "ComputerComponent"
}
