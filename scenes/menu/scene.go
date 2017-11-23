package menu

import (
	"engo.io/ecs"
)

// Scene defines a scene for the main menu
type Scene struct{}

// Type defines the scene
func (s *Scene) Type() string {
	return "MenuScene"
}

// Preload
func (s *Scene) Preload() {
	//
}

// Setup
func (s *Scene) Setup(*ecs.World) {
	//
}
