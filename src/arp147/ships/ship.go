package ships

import (
	"engo.io/ecs"
)

type Ship interface {
	Name() string
	Load()
	Entity() *ecs.Entity
}

func New(s interface{}) interface{} {
	s.(Ship).Load()
	return s
}
