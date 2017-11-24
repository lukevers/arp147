package clock

import (
	"fmt"
	"time"
)

// TODO
type Clock struct {
	Time time.Time
}

// New creates a new global clock at the default starting time.
func New() *Clock {
	return &Clock{
		Time: time.Date(
			6000,
			time.January,
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		),
	}
}

// String returns the current date and time in a string format.
func (c *Clock) String() string {
	return fmt.Sprintf(
		"%d %s, %d - %02d:%02d",
		c.Time.Day(),
		c.Time.Month().String(),
		c.Time.Year(),
		c.Time.Hour(),
		c.Time.Minute(),
	)
}
