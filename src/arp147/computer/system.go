package computer

import (
	"engo.io/ecs"
)

type ComputerSystem struct {
	ecs.LinearSystem
}

func (c *ComputerSystem) Type() string {
	return "ComputerSystem"
}

func (c *ComputerSystem) UpdateEntity(entity *ecs.Entity, dt float32) {
	//
}
