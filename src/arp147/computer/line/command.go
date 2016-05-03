package line

type Command struct {
	Command   string
	Arguments []string
	Directory string
	User      string
	Group     string
}
