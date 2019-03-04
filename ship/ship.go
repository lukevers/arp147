package ship

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type Ship struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	shield *Shield
	engine *Engine
	energy *Energy
	health *Health

	spriteSheet *common.Spritesheet
}

func New() *Ship {
	ship := Ship{BasicEntity: ecs.NewBasic()}

	ship.engine = &Engine{Fuel: 80}
	ship.shield = &Shield{Level: 0}
	ship.energy = &Energy{Solar: 80}
	ship.health = &Health{Level: 100}

	return &ship
}

func (s *Ship) SetPosition(pos engo.Point) {
	pos.X -= s.SpaceComponent.Width / 2
	pos.Y -= s.SpaceComponent.Height / 2
	s.SpaceComponent.Position = pos
}

func (s *Ship) SetSpritesheet(sheet string, width, height float32) {
	s.spriteSheet = common.NewSpritesheetFromFile(sheet, int(width), int(height))

	s.SpaceComponent = common.SpaceComponent{
		Width:  width,
		Height: height,
	}

	s.RenderComponent = common.RenderComponent{
		Drawable: s.spriteSheet.Cell(0),
		Scale:    engo.Point{1, 1},
	}
}

func (s *Ship) AddToWorld(world *ecs.World) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&s.BasicEntity, &s.RenderComponent, &s.SpaceComponent)
		}
	}
}

func (s *Ship) SetSpriteCell(cell int) {
	s.RenderComponent.Drawable = s.spriteSheet.Cell(cell)
}
