package planet

import (
	"testing"
)

func TestPlanetCreate(t *testing.T) {
	p := New(64, TypePlanet)
	p.Generate()
}
