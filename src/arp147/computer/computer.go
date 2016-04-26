package computer

import (
	"arp147/computer/window"
	"arp147/systems/key"
	"arp147/ui/text"
	"engo.io/ecs"
)

type Computer struct {
	Active  bool
	Windows []*window.Window
	world   *ecs.World
	entity  *ecs.Entity
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
		Active:  false,
		Windows: make([]*window.Window, 0),
		world:   world,
	}

	return c
}

func (c *Computer) StartSession() {
	c.Active = true
	if len(c.Windows) < 1 {
		c.Windows = append(c.Windows, window.New(c.world))
	}

	c.world.AddEntity(c.Windows[0].Entity())
	c.Windows[0].StartSession()
}
