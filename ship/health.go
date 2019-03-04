package ship

const (
	HealthLevelMin HealthLevel = 0
	HealthLevelMax HealthLevel = 100
)

type (
	HealthLevel int
)

type Health struct {
	Level HealthLevel
}
