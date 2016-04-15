package computer

import (
	"arp147/events/key"
	"engo.io/engo"
	"util/log"
)

type Computer struct {
	keyListener *key.Listener
}

func New() *Computer {
	c := &Computer{
		keyListener: key.New(),
	}

	c.registerKeys()
	return c
}

func (c *Computer) registerKeys() {
	c.keyListener.On(engo.A, func(key engo.Key) {
		log.Stdout.Println(key)
	})

	c.keyListener.On(engo.B, func(key engo.Key) {
		log.Stdout.Println(key)
	})
}
