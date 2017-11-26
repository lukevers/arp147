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
	mods := &Modifiers{
		Alt:     engo.Input.Button("alt").Down(),
		Control: engo.Input.Button("control").Down(),
		Shift:   engo.Input.Button("shift").Down(),
		Super:   engo.Input.Button("super").Down(),
	}

	for _, key := range *sceneRegistry {
		if btn := engo.Input.Button(key.Name); btn.Down() && key.Down != nil {
			key.Down(mods)
		} else if btn.JustPressed() && key.JustPressed != nil {
			key.JustPressed(mods)
		} else if btn.JustReleased() && key.JustReleased != nil {
			key.JustReleased(mods)
		}
	}
}
