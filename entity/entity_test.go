package entity

import (
	"image"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestEntityCreate(t *testing.T) {
	e := New(SizeViewer, TypePlanet, false)
	e.Generate(TypePlanet)

	saveImage("out.png", e.full)
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
