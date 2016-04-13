package text

import (
	"engo.io/ecs"
)

type Mouse struct {
	Clicked       func(entity *ecs.Entity, dt float32)
	Released      func(entity *ecs.Entity, dt float32)
	Hovered       func(entity *ecs.Entity, dt float32)
	Dragged       func(entity *ecs.Entity, dt float32)
	RightClicked  func(entity *ecs.Entity, dt float32)
	RightReleased func(entity *ecs.Entity, dt float32)
	Enter         func(entity *ecs.Entity, dt float32)
	Leave         func(entity *ecs.Entity, dt float32)
}
