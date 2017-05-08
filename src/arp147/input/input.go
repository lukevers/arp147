package input

import (
	"engo.io/ecs"
	"engo.io/engo"
)

var registry *[]Key

type InputSystem struct {
	entities []inputEntity
}

type inputEntity struct {
	*ecs.BasicEntity
}

type Key struct {
	Name string
	Keys []engo.Key

	Down         func()
	JustPressed  func()
	JustReleased func()
}

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

func (i *InputSystem) Add(basic *ecs.BasicEntity) {
	i.entities = append(i.entities, inputEntity{basic})
}

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
