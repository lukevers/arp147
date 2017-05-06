package engo

import (
	"time"
)

// The amound of nano seconds in a second.
const secondsInNano int64 = 1000000000

// A Clock is a measurement built in `engo` to measure the actual frames per seconds (framerate).
type Clock struct {
	counter   uint32
	perSecond uint32

	deltaStamp int64
	elapsStamp int64
	frameStamp int64
	startStamp int64
}

// NewClock creates a new timer which allows you to measure ticks per seconds. Be sure to call `Tick()` whenever you
// want a tick to occur - it does not automatically tick each frame.
func NewClock() *Clock {
	currStamp := time.Now().UnixNano()

	clock := new(Clock)
	clock.frameStamp = currStamp
	clock.startStamp = currStamp
	return clock
}

// Tick indicates a new tick/frame has occurred.
func (c *Clock) Tick() {
	currStamp := time.Now().UnixNano()

	c.counter += 1

	c.deltaStamp = currStamp - c.frameStamp
	c.frameStamp = currStamp

	c.elapsStamp += c.deltaStamp
	if secondsInNano <= c.elapsStamp {
		c.elapsStamp %= secondsInNano
		c.perSecond = c.counter
		c.counter = 0
	}
}

// Delta is the amount of seconds between the last tick and the one before that
func (c *Clock) Delta() float32 {
	return float32(float64(c.deltaStamp) / float64(secondsInNano))
}

// FPS is the amount of frames per second, computed every time a tick occurs at least a second after the previous update
func (c *Clock) FPS() float32 {
	return float32(c.perSecond)
}

// Time is the number of seconds the clock has been running
func (c *Clock) Time() float32 {
	currStamp := time.Now().UnixNano()
	return float32(float64(currStamp-c.startStamp) / float64(secondsInNano))
}
