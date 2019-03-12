package planet

import (
	"image"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestPlanetCreate(t *testing.T) {
	p := New(SizeViewer, TypePlanet, false)
	p.Generate(TypePlanet)

	saveImage("out.png", p.full)
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
