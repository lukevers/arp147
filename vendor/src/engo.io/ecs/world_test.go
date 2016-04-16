package ecs

import (
	"log"
	"testing"
)

type TestSystem struct {
	LinearSystem
}

func (ts *TestSystem) New(*World) {}
func (*TestSystem) Type() string  { return "TestSystem" }

func (ts *TestSystem) UpdateEntity(e *Entity, dt float32) {}

func TestAddEntity(t *testing.T) {
	world := World{}
	world.New()
	entityOne := NewEntity()
	world.AddEntity(entityOne)
	entityTwo := NewEntity()
	world.AddEntity(entityTwo)
	if len(world.Entities()) != 2 {
		log.Printf("Entities not added.  %d != 2: %+v\n", len(world.Entities()), world.Entities())
		t.Fail()
	}
}

func TestAddSystem(t *testing.T) {
	world := World{}
	world.New()

	before := len(world.Systems())

	system := &TestSystem{}
	world.AddSystem(system)

	if len(world.Systems()) != before+1 {
		t.Fail()
	}
}

func TestAddComponent(t *testing.T) {
	world := World{}
	world.New()
	world.AddSystem(&TestSystem{})
	entity := NewEntity("TestSystem")
	world.AddEntity(entity)
	component := &MyComponent1{5}
	entity.AddComponent(component)
	if len(entity.components) != 1 {
		t.Fail()
	}
}
