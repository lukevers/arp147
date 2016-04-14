package display

import (
	"arp147/scenes/newgame"
	"engo.io/ecs"
	"engo.io/engo"
)

type FakePlayerSystem struct {
	ecs.LinearSystem
}

func (f *FakePlayerSystem) Type() string {
	return "FakePlayerSystem"
}

func (f *FakePlayerSystem) UpdateEntity(entity *ecs.Entity, dt float32) {
	var (
		fc *FakePlayerComponent
		rc *engo.RenderComponent
		sc *engo.SpaceComponent
		mc *engo.MouseComponent
		ok bool
	)

	if fc, ok = entity.ComponentFast(fc).(*FakePlayerComponent); !ok {
		return
	}

	if rc, ok = entity.ComponentFast(rc).(*engo.RenderComponent); !ok {
		return
	}

	if sc, ok = entity.ComponentFast(sc).(*engo.SpaceComponent); !ok {
		return
	}

	if mc, ok = entity.ComponentFast(mc).(*engo.MouseComponent); !ok {
		return
	}

	d := rc.Drawable()

	if mc.Enter {
		rc.SetDrawable(engo.Files.Image("ship-03-full-shields.png"))
	} else if mc.Leave {
		rc.SetDrawable(engo.Files.Image("ship-00-no-shields.png"))
	}

	if !*fc.Begin {
		return
	}

	width := engo.Width()
	height := engo.Height()

	dx := d.Width()
	dy := d.Height()

	x := (width / 2) - (dx / 2)
	y := (height / 2) - (dy / 2)

	if fc.Y == 0 {
		fc.Y = y
	}

	fc.Y = fc.Y - 10
	sc.Position = engo.Point{X: x, Y: fc.Y}

	// Set scene
	if fc.Y < -dy {
		engo.SetScene(&newgame.Scene{}, true)
	}
}
