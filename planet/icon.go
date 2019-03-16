package planet

import (
	"image"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/disintegration/gift"
	"github.com/nfnt/resize"
)

type Icon struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	image       image.Image
	spriteSheet *common.Spritesheet
}

func newIcon(img image.Image, size float64) *Icon {
	i := &Icon{
		image:       img,
		BasicEntity: ecs.NewBasic(),
	}

	i.generate(uint(size))
	i.SetSpritesheet()

	return i
}

func (i *Icon) generate(size uint) {
	i.image = resize.Thumbnail(size, size, i.image, resize.NearestNeighbor)
}

func (i *Icon) SetPosition(pos engo.Point) {
	pos.X -= i.SpaceComponent.Width / 2
	pos.Y -= i.SpaceComponent.Height / 2
	i.SpaceComponent.Position = pos
}

func (i *Icon) SetSpritesheet() {
	g := gift.New()
	dst := image.NewNRGBA(g.Bounds(i.image.Bounds()))
	g.Draw(dst, i.image)

	width := dst.Bounds().Size().X
	height := dst.Bounds().Size().Y

	texture := common.NewTextureResource(common.NewImageObject(dst))
	i.spriteSheet = common.NewSpritesheetFromTexture(
		&texture,
		width,
		height,
	)

	i.SpaceComponent = common.SpaceComponent{
		Width:  float32(width),
		Height: float32(height),
	}

	i.RenderComponent = common.RenderComponent{
		Drawable: i.spriteSheet.Cell(0),
		Scale:    engo.Point{X: 1, Y: 1},
	}
}

func (i *Icon) AddToWorld(world *ecs.World) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&i.BasicEntity, &i.RenderComponent, &i.SpaceComponent)
		}
	}
}
