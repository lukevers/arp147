package terminal

type Command struct {
	Command string
	Args    []string
	Input   Device
	Output  Device
}
