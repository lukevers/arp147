package gerschkin

import (
	"arp147/ships"
	"arp147/ui/position"
	"engo.io/ecs"
	"engo.io/engo"
)

type Ship struct {
	ships.Ship
	Data   ships.Data
	entity *ecs.Entity
}

func (s *Ship) Name() string {
	return "Gerschkin"
}

func (s *Ship) Load() {
	s.Data.Spritesheet = engo.NewSpritesheetFromFile("gerschkin.png", 108, 126)
}

func (s *Ship) Render() {
	texture := s.Data.Spritesheet.Drawable(0)
	render := engo.NewRenderComponent(texture, engo.Point{1, 1}, s.Name())
	s.entity.AddComponent(render)

	pos := position.Position{
		Point:    engo.Point{0, 0},
		Position: position.CENTER_CENTER,
	}

	s.entity.AddComponent(&engo.SpaceComponent{
		Position: pos.Calculate(texture.Width(), texture.Height()),
		Width:    texture.Width(),
		Height:   texture.Height(),
	})
}

func (s *Ship) Entity(systems ...string) *ecs.Entity {
	s.entity = ecs.NewEntity(append(systems, "RenderSystem")...)
	s.Render()
	return s.entity
}
