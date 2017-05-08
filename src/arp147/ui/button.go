package ui

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

// ButtonControlComponent is an extension of common.MouseComponent and contains
// functions for each Mouse event.
type ButtonControlComponent struct {
	Clicked       func(entity *ecs.BasicEntity, dt float32)
	Released      func(entity *ecs.BasicEntity, dt float32)
	Hovered       func(entity *ecs.BasicEntity, dt float32)
	Dragged       func(entity *ecs.BasicEntity, dt float32)
	RightClicked  func(entity *ecs.BasicEntity, dt float32)
	RightReleased func(entity *ecs.BasicEntity, dt float32)
	Enter         func(entity *ecs.BasicEntity, dt float32)
	Leave         func(entity *ecs.BasicEntity, dt float32)
}

// ButtonControlSystem is an extension of common.MouseSystem and listens for
// mouse events and calls functions set on the ButtonControlComponent
// accordingly.
type ButtonControlSystem struct {
	entities []buttonEntity
}

type buttonEntity struct {
	*ecs.BasicEntity
	*common.MouseComponent
	*ButtonControlComponent
}

// OnClicked is a function that is called when an entity with a
// ButtonControlComponent on it.
func (b *ButtonControlComponent) OnClicked(fn func(basic *ecs.BasicEntity, dt float32)) {
	b.Clicked = fn
}

// OnReleased is a function that is called when an entity with a
// ButtonControlComponent on it.
func (b *ButtonControlComponent) OnReleased(fn func(basic *ecs.BasicEntity, dt float32)) {
	b.Released = fn
}

// OnHovered is a function that is called when an entity with a
// ButtonControlComponent on it.
func (b *ButtonControlComponent) OnHovered(fn func(basic *ecs.BasicEntity, dt float32)) {
	b.Hovered = fn
}

// OnDragged is a function that is called when an entity with a
// ButtonControlComponent on it.
func (b *ButtonControlComponent) OnDragged(fn func(basic *ecs.BasicEntity, dt float32)) {
	b.Dragged = fn
}

// OnRightClicked is a function that is called when an entity with a
// ButtonControlComponent on it.
func (b *ButtonControlComponent) OnRightClicked(fn func(basic *ecs.BasicEntity, dt float32)) {
	b.RightClicked = fn
}

// OnRightReleased is a function that is called when an entity with a
// ButtonControlComponent on it.
func (b *ButtonControlComponent) OnRightReleased(fn func(basic *ecs.BasicEntity, dt float32)) {
	b.RightReleased = fn
}

// OnEnter is a function that is called when an entity with a
// ButtonControlComponent on it.
func (b *ButtonControlComponent) OnEnter(fn func(basic *ecs.BasicEntity, dt float32)) {
	b.Enter = fn
}

// OnLease is a function that is called when an entity with a
// ButtonControlComponent on it.
func (b *ButtonControlComponent) OnLease(fn func(basic *ecs.BasicEntity, dt float32)) {
	b.Leave = fn
}

func (b *ButtonControlSystem) Update(dt float32) {
	for _, e := range b.entities {
		if e.MouseComponent.Clicked && e.ButtonControlComponent.Clicked != nil {
			e.ButtonControlComponent.Enter(e.BasicEntity, dt)
		} else if e.MouseComponent.Released && e.ButtonControlComponent.Released != nil {
			e.ButtonControlComponent.Released(e.BasicEntity, dt)
		} else if e.MouseComponent.Hovered && e.ButtonControlComponent.Hovered != nil {
			e.ButtonControlComponent.Hovered(e.BasicEntity, dt)
		} else if e.MouseComponent.Dragged && e.ButtonControlComponent.Dragged != nil {
			e.ButtonControlComponent.Dragged(e.BasicEntity, dt)
		} else if e.MouseComponent.RightClicked && e.ButtonControlComponent.RightClicked != nil {
			e.ButtonControlComponent.RightClicked(e.BasicEntity, dt)
		} else if e.MouseComponent.RightReleased && e.ButtonControlComponent.RightReleased != nil {
			e.ButtonControlComponent.RightReleased(e.BasicEntity, dt)
		} else if e.MouseComponent.Enter && e.ButtonControlComponent.Enter != nil {
			e.ButtonControlComponent.Enter(e.BasicEntity, dt)
		} else if e.MouseComponent.Leave && e.ButtonControlComponent.Leave != nil {
			e.ButtonControlComponent.Leave(e.BasicEntity, dt)
		}
	}
}

func (b *ButtonControlSystem) Add(basic *ecs.BasicEntity, mouse *common.MouseComponent, button *ButtonControlComponent) {
	b.entities = append(b.entities, buttonEntity{basic, mouse, button})
}

func (b *ButtonControlSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range b.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		b.entities = append(b.entities[:delete], b.entities[delete+1:]...)
	}
}
