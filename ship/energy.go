package ship

const (
	EnergyLevelMin EnergyLevel = 0
	EnergyLevelMax EnergyLevel = 100
)

type (
	EnergyLevel int
)

type Energy struct {
	Solar EnergyLevel
}
