package ship

import (
	"fmt"
)

func (s *Ship) HudShieldLevel() string {
	return fmt.Sprintf(
		"%d/%d - SHIELD STRENGTH",
		s.Shield.Level,
		SheldLevelMax,
	)
}

func (s *Ship) HudShipHealth() string {
	return fmt.Sprintf(
		"%d/%d - SHIP HEALTH",
		s.Health.Level,
		HealthLevelMax,
	)
}

func (s *Ship) HudEnergySolarLevel() string {
	return fmt.Sprintf(
		"SOLAR LEVEL - %d/%d",
		s.Energy.Solar,
		EnergyLevelMax,
	)
}

func (s *Ship) HudFuelLevel() string {
	return fmt.Sprintf(
		"FUEL  LEVEL - %d/%d",
		s.Engine.Fuel,
		int(FuelLevelMax),
	)
}
