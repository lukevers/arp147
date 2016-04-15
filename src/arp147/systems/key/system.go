package key

import (
	"engo.io/ecs"
	"engo.io/engo"
)

type KeySystem struct {
	ecs.LinearSystem
}

func (k *KeySystem) Type() string {
	return "KeySystem"
}

func (k *KeySystem) UpdateEntity(entity *ecs.Entity, dt float32) {
	var (
		kc *KeyComponent
		ok bool
	)

	if kc, ok = entity.ComponentFast(kc).(*KeyComponent); !ok {
		return
	}

	for key, fns := range kc.keys {
		if engo.Keys.Get(key).JustPressed() {
			for _, fn := range fns {
				fn(key)
			}
		}
	}
}
