package planet

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/disintegration/gift"
	"github.com/fogleman/gg"
)

type Type int

const (
	TypePlanet Type = iota
	TypeMoon
	TypeStar
)

const (
	SizeViewer = 64
)

type Planet struct {
	seed int64
	size float64

	full  image.Image
	main  image.Image
	moons []image.Image

	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	spriteSheet *common.Spritesheet
}

func New(size float64, t Type, initalize bool) *Planet {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	p := &Planet{
		seed: seed,
		size: size,
	}

	if initalize {
		p.Generate(t)
		p.SetSpritesheet()
	}

	return p
}

func randInt(min, max int) uint8 {
	return uint8(min + rand.Intn(max-min))
}

func (p *Planet) Generate(t Type) {
	p.main = p.generate(p.size, t)

	if t == TypePlanet {
		for i := 0; i < int(randInt(0, 4)); i++ {
			moon := p.generate(float64(randInt(int(p.size/16), int(p.size/4))), TypeMoon)
			p.moons = append(p.moons, moon)
		}
	}

	p.full = p.patchMoons()
}

func (p *Planet) generate(size float64, t Type) image.Image {
	dc := gg.NewContext(int(size*2), int(size*2))
	dc.DrawCircle(size, size, size)

	grad := gg.NewLinearGradient(20, 320, 400, 20)
	grad.AddColorStop(0, color.RGBA{randInt(0, 255), randInt(0, 255), randInt(0, 255), 255})
	grad.AddColorStop(1, color.RGBA{randInt(0, 255), randInt(0, 255), randInt(0, 255), 255})
	grad.AddColorStop(.3, color.RGBA{randInt(0, 255), randInt(0, 255), randInt(0, 255), 255})
	grad.AddColorStop(.5, color.RGBA{randInt(0, 255), randInt(0, 255), randInt(0, 255), 255})

	dc.SetFillStyle(grad)
	dc.Fill()

	img := dc.Image()

	for _, filter := range p.filters(t, size) {
		g := gift.New(filter)
		dst := image.NewNRGBA(g.Bounds(img.Bounds()))
		g.Draw(dst, img)
		img = dst
	}

	return img
}

// for testing
func saveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("png.Encode failed: %v", err)
	}
}

func (p *Planet) filters(t Type, size float64) map[string]gift.Filter {
	// Make changes before applying filters if TypeMoon
	if t == TypeMoon {
		// Double the "size" when applying filters related to sizing to
		// increase the pixelate filter.
		size *= 2
	}

	return map[string]gift.Filter{
		"color_balance": gift.ColorBalance(3, 5, -10),
		"color_func": gift.ColorFunc(
			func(r0, g0, b0, a0 float32) (r, g, b, a float32) {
				r = 1 - r0
				g = g0 + 0.5
				b = b0 + .3
				a = a0
				return r, g, b, a
			},
		),
		"pixelate":     gift.Pixelate(int(size / 10.0)),
		"unsharp_mask": gift.UnsharpMask(1, 32, 0),
	}
}

func (p *Planet) patchMoons() image.Image {
	size := int(p.size * 4)
	dc := gg.NewContext(size, size)
	dc.DrawImageAnchored(p.main, p.main.Bounds().Size().X, p.main.Bounds().Size().Y, .5, .5)

	var x, y int
	var ax, ay float64

	for i, moon := range p.moons {
		xoffset := int(p.size/4) + moon.Bounds().Size().X
		yoffset := int(p.size/4) + moon.Bounds().Size().Y

		if moon.Bounds().Size().X <= 12 {
			xoffset *= 2
			yoffset *= 2
		}

		xoffset -= int(randInt(0, int(p.size/6)))
		yoffset -= int(randInt(0, int(p.size/6)))

		switch i {
		case 0:
			x = 0 + xoffset
			y = 0 + yoffset
			ax = 0
			ay = 0
		case 1:
			x = size - xoffset
			y = 0 + yoffset
			ax = 1
			ay = 0
		case 2:
			x = 0 + xoffset
			y = size - yoffset
			ax = 0
			ay = 1

		case 3:
			x = size - xoffset
			y = size - yoffset
			ax = 1
			ay = 1
		}

		dc.DrawImageAnchored(moon, x, y, ax, ay)
	}

	return dc.Image()
}

func (p *Planet) AddToWorld(world *ecs.World) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&p.BasicEntity, &p.RenderComponent, &p.SpaceComponent)
		}
	}

}

func (p *Planet) SetPosition(pos engo.Point) {
	pos.X -= p.SpaceComponent.Width / 2
	pos.Y -= p.SpaceComponent.Height / 2
	p.SpaceComponent.Position = pos
}

func (p *Planet) SetSpritesheet() {
	g := gift.New()
	dst := image.NewNRGBA(g.Bounds(p.full.Bounds()))
	g.Draw(dst, p.full)

	width := dst.Bounds().Size().X
	height := dst.Bounds().Size().Y

	texture := common.NewTextureResource(common.NewImageObject(dst))
	p.spriteSheet = common.NewSpritesheetFromTexture(
		&texture,
		width,
		height,
	)

	p.SpaceComponent = common.SpaceComponent{
		Width:  float32(width),
		Height: float32(height),
	}

	p.RenderComponent = common.RenderComponent{
		Drawable: p.spriteSheet.Cell(0),
		Scale:    engo.Point{1, 1},
	}
}
