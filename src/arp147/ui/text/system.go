package text

import (
	"engo.io/ecs"
	"engo.io/engo"
)

type TextControlSystem struct {
	ecs.LinearSystem
}

func (t *TextControlSystem) New(w *ecs.World) {
	// ...
}

func (t *TextControlSystem) Type() string {
	return "TextControlSystem"
}

// Priority defines the order of which systems to load first. The higher the
// priority, the higher up in the loop the system is loaded. The MouseSystem
// has a priority of 10, and the TextControlSystem needs on the MouseSystem.
func (t *TextControlSystem) Priority() int {
	return 10
}

func (t *TextControlSystem) UpdateEntity(entity *ecs.Entity, dt float32) {
	var (
		mc *engo.MouseComponent
		tc *TextControlComponent
		ok bool
	)

	if mc, ok = entity.ComponentFast(mc).(*engo.MouseComponent); !ok {
		return
	}

	if tc, ok = entity.ComponentFast(tc).(*TextControlComponent); !ok {
		return
	}

	if mc.Clicked && tc.Mouse.Clicked != nil {
		tc.Mouse.Clicked(entity, dt)
	} else if mc.Released && tc.Mouse.Released != nil {
		tc.Mouse.Released(entity, dt)
	} else if mc.Hovered && tc.Mouse.Hovered != nil {
		tc.Mouse.Hovered(entity, dt)
	} else if mc.Dragged && tc.Mouse.Dragged != nil {
		tc.Mouse.Dragged(entity, dt)
	} else if mc.RightClicked && tc.Mouse.RightClicked != nil {
		tc.Mouse.RightClicked(entity, dt)
	} else if mc.RightReleased && tc.Mouse.RightReleased != nil {
		tc.Mouse.RightReleased(entity, dt)
	} else if mc.Enter && tc.Mouse.Enter != nil {
		tc.Mouse.Enter(entity, dt)
	} else if mc.Leave && tc.Mouse.Leave != nil {
		tc.Mouse.Leave(entity, dt)
	}
}
