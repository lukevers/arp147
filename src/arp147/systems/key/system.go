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
				fn(key, caps())
			}
		}
	}
}

func caps() bool {
	if engo.Keys.Get(engo.LeftShift).Down() {
		return true
	} else if engo.Keys.Get(engo.RightShift).Down() {
		return true
	} else if engo.Keys.Get(engo.CapsLock).Down() {
		return true
	} else {
		return false
	}
}
