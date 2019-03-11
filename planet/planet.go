package planet

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/disintegration/gift"
	"github.com/fogleman/gg"
)

type Type int

const (
	TypePlanet Type = iota
	TypeMoon
	TypeStar
)

type Planet struct {
	seed int64
	size float64

	main  image.Image
	moons []image.Image
}

func New(size float64) *Planet {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	return &Planet{
		seed: seed,
		size: size,
	}
}

func randInt(min, max int) uint8 {
	return uint8(min + rand.Intn(max-min))
}

func (p *Planet) Generate(t Type) {
	p.main = p.generate(p.size, t)
	saveImage("out.png", p.main)

	if t == TypePlanet {
		for i := 0; i < int(randInt(0, 5)); i++ {
			moon := p.generate(float64(randInt(int(p.size/16), int(p.size/4))), TypeMoon)
			p.moons = append(p.moons, moon)

			saveImage(fmt.Sprintf("moon-%d.png", i), moon)
		}
	}
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
	// TODO: different set of filters per Type

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
		"unsharp_mask": gift.UnsharpMask(1, 20, 0),
	}
}
