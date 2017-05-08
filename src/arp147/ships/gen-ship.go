// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package ships

import (
	"arp147/logging"
	"arp147/ui"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// TheGerschkin is a ship.
type TheGerschkin struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	Position    ui.Position
	Spritesheet *common.Spritesheet
}

// We only want to load the spritesheet once, so we're storing a per-ship bool
// and will use that to check if we should load it or not each time we create a
// new instance of The Gerschkin ship.
var loadedGerschkinSpritesheet bool = false

// NewGerschkin creates a new instance of TheGerschkin and sets up everything
// it needs to be used.
func NewGerschkin() *TheGerschkin {
	ship := &TheGerschkin{}

	if !loadedGerschkinSpritesheet {
		// If we can't load the texture file, everything will be fucked up and
		// we should just stop.
		if err := engo.Files.Load(TheGerschkinSpritesheetTexturePath); err != nil {
			logging.Stderr.Fatal("Could not load file: ", err)
		}

		loadedGerschkinSpritesheet = true
	}

	// Load the spritesheet.
	ship.Spritesheet = common.NewSpritesheetFromFile(
		TheGerschkinSpritesheetTexturePath,
		TheGerschkinSpritesheetWidth,
		TheGerschkinSpritesheetHeight,
	)

	return ship
}

// AddToWorld adds TheGerschkin ship to the world passed to it.
func (ship *TheGerschkin) AddToWorld(world *ecs.World) {
	ship.BasicEntity = ecs.NewBasic()

	texture := ship.Spritesheet.Drawable(0)
	w := texture.Width()
	h := texture.Height()

	ship.SpaceComponent = common.SpaceComponent{
		Position: ship.Position.Calculate(w, h),
		Width:    w,
		Height:   h,
	}

	ship.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{1, 1},
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&ship.BasicEntity,
				&ship.RenderComponent,
				&ship.SpaceComponent,
			)
		}
	}
}