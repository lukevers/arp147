package input

import (
	"engo.io/ecs"
	"engo.io/engo"
)

// InputSystem listens for key events.
type InputSystem struct {
	entities []inputEntity
}

type inputEntity struct {
	*ecs.BasicEntity
}

// Add takes an entity and adds it to the system
func (i *InputSystem) Add(basic *ecs.BasicEntity) {
	i.entities = append(i.entities, inputEntity{basic})
}

// Remove takes an entity and removes it from the system
func (i *InputSystem) Remove(basic ecs.BasicEntity) {
	delete := -1

	for index, e := range i.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}

	if delete >= 0 {
		i.entities = append(i.entities[:delete], i.entities[delete+1:]...)
	}
}

// Update is called on each frame when the system is in use.
func (i *InputSystem) Update(dt float32) {
	for _, key := range *sceneRegistry {
		if btn := engo.Input.Button(key.Name); btn.JustPressed() && key.OnPress != nil {
			key.OnPress(btn.Triggers[0], i.modifiers())
		} else if btn.JustReleased() && key.OnRelease != nil {
			key.OnRelease(btn.Triggers[0], i.modifiers())
		}
	}
}

func (i *InputSystem) modifiers() *Modifiers {
	return &Modifiers{
		Alt:     engo.Input.Button(ModifierAlt).Down(),
		Control: engo.Input.Button(ModifierControl).Down(),
		Shift:   engo.Input.Button(ModifierShift).Down(),
		Super:   engo.Input.Button(ModifierSuper).Down(),
	}
}
