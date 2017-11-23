package clock

import (
	"fmt"
	"time"
)

// Time contains the global time for the world.
var Time *time.Time

// New creates a new global clock at the default starting time.
func New() {
	t := time.Date(
		6000,
		time.January,
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	Time = &t
}

// String returns the current date and time in a string format.
func String() string {
	return fmt.Sprintf(
		"%d %s, %d - %02d:%02d",
		Time.Day(),
		Time.Month().String(),
		Time.Year(),
		Time.Hour(),
		Time.Minute(),
	)
}
