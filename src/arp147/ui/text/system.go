package text

import (
	"github.com/engoengine/ecs"
	"github.com/engoengine/engo"
)

type TextControlSystem struct {
	ecs.LinearSystem

	Clicked       func(entity *ecs.Entity, dt float32)
	Released      func(entity *ecs.Entity, dt float32)
	Hovered       func(entity *ecs.Entity, dt float32)
	Dragged       func(entity *ecs.Entity, dt float32)
	RightClicked  func(entity *ecs.Entity, dt float32)
	RightReleased func(entity *ecs.Entity, dt float32)
	Enter         func(entity *ecs.Entity, dt float32)
	Leave         func(entity *ecs.Entity, dt float32)
}

func (t *TextControlSystem) Type() string {
	return "TextControlSystem"
}

func (t *TextControlSystem) New(w *ecs.World) {
	// ...
}

func (t *TextControlSystem) UpdateEntity(entity *ecs.Entity, dt float32) {
	var mouse *engo.MouseComponent
	if !entity.Component(&mouse) {
		return
	}

	if mouse.Clicked && t.Clicked != nil {
		t.Clicked(entity, dt)
	} else if mouse.Released && t.Released != nil {
		t.Released(entity, dt)
	} else if mouse.Hovered && t.Hovered != nil {
		t.Hovered(entity, dt)
	} else if mouse.Dragged && t.Dragged != nil {
		t.Dragged(entity, dt)
	} else if mouse.RightClicked && t.RightClicked != nil {
		t.RightClicked(entity, dt)
	} else if mouse.RightReleased && t.RightReleased != nil {
		t.RightReleased(entity, dt)
	} else if mouse.Enter && t.Enter != nil {
		t.Enter(entity, dt)
	} else if mouse.Leave && t.Leave != nil {
		t.Leave(entity, dt)
	}
}
