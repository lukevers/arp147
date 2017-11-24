package clock

import (
	"engo.io/ecs"
	"time"
)

// ClockSystem updates the clock every frame.
type ClockSystem struct {
	Clock *Clock

	entities []clockEntity
	dt       float32
}

type clockEntity struct {
	*ecs.BasicEntity
}

// Add takes an entity and adds it to the system
func (c *ClockSystem) Add(basic *ecs.BasicEntity) {
	c.entities = append(c.entities, clockEntity{basic})
}

// Remove takes an entity and removes it from the system
func (c *ClockSystem) Remove(basic ecs.BasicEntity) {
	delete := -1

	for index, e := range c.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}

	if delete >= 0 {
		c.entities = append(c.entities[:delete], c.entities[delete+1:]...)
	}
}

// Update is called on each frame when the system is in use.
func (c *ClockSystem) Update(dt float32) {
	c.dt += dt
	if c.dt >= 1 {
		c.dt = 0

		// Advance the clock by a minute
		c.Clock.Time = c.Clock.Time.Add(time.Minute * 1)
	}
}
