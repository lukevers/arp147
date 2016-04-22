package computer

import (
	"arp147/systems/key"
	"arp147/ui/text"
	"engo.io/ecs"
)

type Computer struct {
	world  *ecs.World
	entity *ecs.Entity
	lines  map[int]*line
	line   int
}

type line struct {
	text   []*text.Text
	locked bool
}

func New(world *ecs.World) *Computer {
	// Check if the systems we want are already created
	k, t := true, true
	for _, system := range world.Systems() {
		switch system.Type() {
		case "KeySystem":
			k = false
		case "TextControlSystem":
			t = false
		}
	}

	// Add the key system if it's not already created
	if k {
		world.AddSystem(&key.KeySystem{})
	}

	// Add the text control system if it's not already created
	if t {
		world.AddSystem(&text.TextControlSystem{})
	}

	c := &Computer{
		world:  world,
		lines:  make(map[int]*line),
		active: true,
	}

	c.createUI()
	return c
}
