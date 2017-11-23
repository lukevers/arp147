package ui

import (
	"arp147/logging"
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// TODO
type Background struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	Texture string
}

// TODO
type BackgroundSystem struct {
	entities []backgroundEntity
	dt       float32
	offset   float32
	world    *ecs.World
}

type backgroundEntity struct {
	*Background
	*ecs.World
}

// TODO
func NewBackground(background Background) *Background {
	background.BasicEntity = ecs.NewBasic()
	background.RenderComponent = common.RenderComponent{}
	background.SpaceComponent = common.SpaceComponent{}

	// TODO

	return &background
}

// TODO
func (b *Background) Render() {
	// TODO
}

// TODO
func (b *Background) AddToWorld(world *ecs.World) {
	b.Render()

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *BackgroundSystem:
			sys.world = world
			sys.Add(backgroundEntity{b, world})
		}
	}
}

// Add takes an entity and adds it to the system
func (b *BackgroundSystem) Add(background backgroundEntity) {
	b.entities = append(b.entities, background)
}

// Remove takes an entity and removes it from the system
func (b *BackgroundSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range b.entities {
		if e.Background.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		b.entities = append(b.entities[:delete], b.entities[delete+1:]...)
	}
}

// Update is called on each frame when the system is in use.
func (b *BackgroundSystem) Update(dt float32) {
	b.dt += dt
	if b.dt >= 1 {
		b.dt = 0
		b.offset++

		TileWorld(b.world, "textures/space.png", 5*b.offset)
	}
}

// --- old, to be deleted soon
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

type background struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// TileWorld is a helper function that tiles an image on the entire screen.
func TileWorld(world *ecs.World, image string, offset float32) {
	texture, err := common.LoadedSprite(image)
	if err != nil {
		logging.Stderr.Fatal("Could not use loaded sprite: ", err)
	}

	width := engo.GameWidth()
	height := engo.GameHeight()

	w := texture.Width()
	h := texture.Height()

	var x, y float32 = offset, 0

	for {
		var x2 float32 = x
		var y2 float32 = y

		bkg := background{
			BasicEntity: ecs.NewBasic(),
			SpaceComponent: common.SpaceComponent{
				Position: engo.Point{x2, y2},
				Width:    w,
				Height:   h,
			},
			RenderComponent: common.RenderComponent{
				Drawable: texture,
				Scale:    engo.Point{1, 1},
			},
		}

		for _, system := range world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(
					&bkg.BasicEntity,
					&bkg.RenderComponent,
					&bkg.SpaceComponent,
				)
			}
		}

		if x > width {
			x = offset
			y += h
		} else {
			x += w
		}

		if y > height {
			break
		}
	}
}
