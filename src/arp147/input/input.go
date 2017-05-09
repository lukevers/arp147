package input

import (
	"engo.io/ecs"
	"engo.io/engo"
)

var registry *[]Key

// InputSystem listens for key events.
type InputSystem struct {
	entities []inputEntity
}

type inputEntity struct {
	*ecs.BasicEntity
}

// Key contains the information needed for key events. It also contains three
// functions that can be set on each type of key event.
type Key struct {
	Name string
	Keys []engo.Key

	Down         func()
	JustPressed  func()
	JustReleased func()
}

// RegisterButtons registers buttons with functions for the scene.
func RegisterButtons(buttons []Key) {
	// Store the button registry globally
	registry = &buttons

	// Register all of the buttons
	for _, key := range buttons {
		engo.Input.RegisterButton(
			key.Name,
			key.Keys...,
		)
	}
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
	for _, key := range *registry {
		if btn := engo.Input.Button(key.Name); btn.Down() && key.Down != nil {
			key.Down()
		} else if btn.JustPressed() && key.JustPressed != nil {
			key.JustPressed()
		} else if btn.JustReleased() && key.JustReleased != nil {
			key.JustReleased()
		}
	}
}
