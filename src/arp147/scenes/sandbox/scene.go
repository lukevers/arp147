package sandbox

import (
	"arp147/ui"
	"engo.io/ecs"
	"engo.io/engo/common"
)

type SandboxScene struct {
	//
}

func (*SandboxScene) Type() string {
	return "SandboxScene"
}

func (scene *SandboxScene) Preload() {
	//
}

func (scene *SandboxScene) Setup(world *ecs.World) {
	// Add systems needed for the scene.
	world.AddSystem(&common.RenderSystem{})

	// Tile the world background.
	ui.TileWorld(world, "textures/space.png")

	// Setup supported input for the scene
	scene.SetupInput(world)
}
