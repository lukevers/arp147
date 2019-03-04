package ship

const (
	EnginePowerMin EnginePower = 0
	EnginePowerMax EnginePower = 5

	FuelLevelMin FuelLevel = 0
	FuelLevelMax FuelLevel = 100
)

type (
	EnginePower int
	FuelLevel   int
)

type Engine struct {
	Power EnginePower
	Fuel  FuelLevel
}
