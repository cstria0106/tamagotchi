package game

import (
	"errors"
	"golang.org/x/exp/rand"
	"sync"
	"time"
)

type Game interface {
	Start() error
	Update() error
	WithSystems(f func(systems []System))
	WithComponents(componentIds []ComponentID, f func(components []*Component))
	NewEntity(options ...func(*Entity)) *Entity
	AddEntities(entity ...*Entity)
	RemoveEntity(entity *Entity) bool
	AddSystem(systems ...System) error
}

type BaseGame struct {
	lock     sync.Mutex
	started  bool
	entities []*Entity
	systems  []System
}

func (g *BaseGame) Start() error {
	g.lock.Lock()
	if g.started {
		return errors.New("engine already started")
	}

	g.started = true
	g.lock.Unlock()

	for _, system := range g.systems {
		if system.Init != nil {
			if err := system.Init(g); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *BaseGame) Update() error {
	for _, system := range g.systems {
		if system.Update != nil {
			if err := system.Update(g); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *BaseGame) WithSystems(f func(systems []System)) {
	f(g.systems)
}

func (g *BaseGame) WithComponents(componentIds []ComponentID, f func(components []*Component)) {
	for _, e := range g.entities {
		components, all := e.Components(componentIds)

		if all {
			f(components)
		}
	}
}

func (g *BaseGame) NewEntity(options ...func(*Entity)) *Entity {
	return NewEntity(options...)
}

func (g *BaseGame) AddEntities(entity ...*Entity) {
	g.entities = append(g.entities, entity...)
}

func (g *BaseGame) RemoveEntity(entity *Entity) bool {
	for i, e := range g.entities {
		if entity == e {
			g.entities = append(g.entities[:i], g.entities[i+1:]...)
			return true
		}
	}

	return false
}

func (g *BaseGame) AddSystem(systems ...System) error {
	g.systems = append(g.systems, systems...)

	if g.started {
		for _, system := range systems {
			if system.Init != nil {
				if err := system.Init(g); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func New() (Game, error) {
	rand.Seed(uint64(time.Now().Unix()))

	g := &BaseGame{
		entities: make([]*Entity, 0, 1024),
		systems:  make([]System, 0, 1024),
	}

	return g, nil
}
