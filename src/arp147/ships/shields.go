package ships

const (
	SHIELDS_LEVEL_OFF = 1 << iota
	SHIELDS_LEVEL_LOW
	SHIELDS_LEVEL_PARTIAL
	SHIELDS_LEVEL_HIGH
)

type Shields struct {
	Level int
}
