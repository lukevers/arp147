package ships

import (
	"engo.io/ecs"
)

type Ship interface {
	// The Name func contains the name of the ship
	Name() string

	// The Load func should do any pre-processing or setup before we begin to
	// render and create the Entity.
	Load()

	// The Entity func contains an optional array of systems to include when
	// creating the entity. This is useful for adding certain systems to
	// entities in specific scenes but not globally.
	Entity(...string) *ecs.Entity
}

func New(s interface{}) interface{} {
	s.(Ship).Load()
	return s
}
