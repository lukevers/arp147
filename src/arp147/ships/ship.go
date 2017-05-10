package ships

import (
	"arp147/logging"
	"arp147/ui"
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Generic=Gerschkin"
type Generic generic.Type

// TheGeneric is a ship.
type TheGeneric struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	Position    ui.Position
	Spritesheet *common.Spritesheet
}

// We only want to load the spritesheet once, so we're storing a per-ship bool
// and will use that to check if we should load it or not each time we create a
// new instance of The Generic ship.
var loadedGenericSpritesheet bool = false

// NewGeneric creates a new instance of TheGeneric and sets up everything
// it needs to be used.
func NewGeneric() *TheGeneric {
	ship := &TheGeneric{}

	if !loadedGenericSpritesheet {
		// If we can't load the texture file, everything will be fucked up and
		// we should just stop.
		if err := engo.Files.Load(TheGenericSpritesheetTexturePath); err != nil {
			logging.Stderr.Fatal("Could not load file: ", err)
		}

		loadedGenericSpritesheet = true
	}

	// Load the spritesheet.
	ship.Spritesheet = common.NewSpritesheetFromFile(
		TheGenericSpritesheetTexturePath,
		TheGerschkinSpritesheetWidth,
		TheGerschkinSpritesheetHeight,
	)

	return ship
}

// AddToWorld adds TheGeneric ship to the world passed to it.
func (ship *TheGeneric) AddToWorld(world *ecs.World) {
	ship.BasicEntity = ecs.NewBasic()

	texture := ship.Spritesheet.Drawable(0)
	w := texture.Width() * TheGenericSpritesheetScale.X
	h := texture.Height() * TheGenericSpritesheetScale.Y

	ship.SpaceComponent = common.SpaceComponent{
		Position: ship.Position.Calculate(w, h),
		Width:    w,
		Height:   h,
	}

	ship.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    TheGenericSpritesheetScale,
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
