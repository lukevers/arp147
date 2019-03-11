package planet

import (
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
	TypeSun
	TypeMoon
)

type Planet struct {
	seed int64
	size float64
}

func New(size float64, t Type) *Planet {
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

func (p *Planet) Generate() {
	dc := gg.NewContext(int(p.size*2), int(p.size*2))
	dc.DrawCircle(p.size, p.size, p.size)

	grad := gg.NewLinearGradient(20, 320, 400, 20)
	grad.AddColorStop(0, color.RGBA{randInt(0, 255), randInt(0, 255), randInt(0, 255), 255})
	grad.AddColorStop(1, color.RGBA{randInt(0, 255), randInt(0, 255), randInt(0, 255), 255})
	grad.AddColorStop(.3, color.RGBA{randInt(0, 255), randInt(0, 255), randInt(0, 255), 255})
	grad.AddColorStop(.5, color.RGBA{randInt(0, 255), randInt(0, 255), randInt(0, 255), 255})

	dc.SetFillStyle(grad)
	dc.Fill()

	dc.SavePNG("out.png")

	filters := map[string]gift.Filter{
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
		"pixelate":     gift.Pixelate(int(p.size / 10)),
		"unsharp_mask": gift.UnsharpMask(1, 20, 0),
	}

	for _, filter := range filters {
		src := loadImage("out.png")
		g := gift.New(filter)
		dst := image.NewNRGBA(g.Bounds(src.Bounds()))
		g.Draw(dst, src)
		saveImage("out.png", dst)
	}

}

func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
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
