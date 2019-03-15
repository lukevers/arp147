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

	Shield *Shield
	Engine *Engine
	Energy *Energy
	Health *Health

	spriteSheet *common.Spritesheet
	defined     Defined
}

func New(d Defined) *Ship {
	s := Ship{
		BasicEntity: ecs.NewBasic(),
		Engine:      &Engine{Fuel: 80},
		Shield:      &Shield{Level: 0},
		Energy:      &Energy{Solar: 80},
		Health:      &Health{Level: 100},
		defined:     d,
	}

	texture, width, height := s.defined.GetSpritesheet()
	s.spriteSheet = common.NewSpritesheetFromFile(texture, width, height)

	s.SpaceComponent = common.SpaceComponent{
		Width:  float32(width),
		Height: float32(height),
	}

	s.RenderComponent = common.RenderComponent{
		Drawable: s.spriteSheet.Cell(0),
		Scale:    engo.Point{1, 1},
	}

	ship := &s

	engo.Mailbox.Dispatch(NewShipMessage{
		Ship: ship,
	})

	return ship
}

func (s *Ship) SetPosition(pos engo.Point) {
	pos.X -= s.SpaceComponent.Width / 2
	pos.Y -= s.SpaceComponent.Height / 2
	s.SpaceComponent.Position = pos
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
