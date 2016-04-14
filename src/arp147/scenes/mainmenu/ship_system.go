package mainmenu

import (
	"arp147/scenes/newgame"
	"engo.io/ecs"
	"engo.io/engo"
)

type ShipSystem struct {
	ecs.LinearSystem
}

func (s *ShipSystem) Type() string {
	return "ShipSystem"
}

func (s *ShipSystem) UpdateEntity(entity *ecs.Entity, dt float32) {
	var (
		rc *engo.RenderComponent
		sp *engo.SpaceComponent
		sc *ShipComponent
		ok bool
	)

	if rc, ok = entity.ComponentFast(rc).(*engo.RenderComponent); !ok {
		return
	}

	if sp, ok = entity.ComponentFast(sp).(*engo.SpaceComponent); !ok {
		return
	}

	if sc, ok = entity.ComponentFast(sc).(*ShipComponent); !ok {
		return
	}

	if !*sc.Begin {
		return
	}

	var mod float32 = -0.05
	if rc.Scale().X+mod <= 0 {
		engo.SetScene(&newgame.Scene{}, true)
		return
	} else {
		// Scale
		s := rc.Scale()
		s.AddScalar(mod)
		rc.SetScale(s)

		// Keep centered
		dx := sp.Width
		dy := sp.Height
		width := engo.Width()
		height := engo.Height()
		sp.Position = engo.Point{
			X: (width / 2) - (dx * s.X / 2),
			Y: (height / 2) - (dy * s.Y / 2),
		}
	}
}
