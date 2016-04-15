package key

import (
	"engo.io/engo"
)

type Listener struct {
	keys map[engo.Key][]func(engo.Key)
}

func New() *Listener {
	l := &Listener{
		keys: make(map[engo.Key][]func(engo.Key)),
	}

	go l.listen()
	return l
}

func (l *Listener) On(key engo.Key, fn func(engo.Key)) {
	l.keys[key] = append(l.keys[key], fn)
}

func (l *Listener) listen() {
	clock := engo.Time
	delta := clock.Delta()

	for {
		for key, fns := range l.keys {
			k := engo.Keys.Get(key)
			if k.State() == engo.KEY_STATE_JUST_DOWN {
				if clock.Delta() == delta {
					continue
				}

				for _, fn := range fns {
					fn(key)
				}
			}
		}

		delta = clock.Delta()
	}
}
