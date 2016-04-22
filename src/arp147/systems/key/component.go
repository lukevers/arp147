package key

import (
	"engo.io/ecs"
	"engo.io/engo"
)

type KeyComponent struct {
	keys map[engo.Key][]func(engo.Key, bool)
}

func (k *KeyComponent) New(world *ecs.World) {
	k.keys = make(map[engo.Key][]func(engo.Key, bool))
}

func (k *KeyComponent) Type() string {
	return "KeyComponent"
}

func (k *KeyComponent) On(key engo.Key, fn func(engo.Key, bool)) {
	if k.keys == nil {
		k.New(nil)
	}

	k.keys[key] = append(k.keys[key], fn)
}
