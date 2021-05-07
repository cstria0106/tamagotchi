package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/segmentio/ksuid"
	"sync"
)

type Entity interface {
	AddComponent(c Component) error
	GetComponent(componentUID ksuid.KSUID) Component
	GetComponents(componentUID ksuid.KSUID) []Component
	RemoveComponent(component Component) bool
	Clone() Entity
	Init() error
	Dispose()
	Update() error
	Draw(screen *ebiten.Image)

	Game() *Game
	ID() string

	assignComponents(components ...Component)
}

type baseEntity struct {
	game        *Game
	id          string
	mutex       sync.Mutex
	initialized bool
	components  []Component
}

func (g *Game) NewBaseEntity(options ...func(Entity)) *baseEntity {
	return NewBaseEntity(g, options...)
}

func NewBaseEntity(g *Game, options ...func(Entity)) *baseEntity {
	entity := baseEntity{
		game:        g,
		id:          ksuid.New().String(),
		mutex:       sync.Mutex{},
		initialized: false,
		components:  make([]Component, 0),
	}

	for _, option := range options {
		option(&entity)
	}

	return &entity
}

func NewBaseEntityComponents(components ...Component) func(Entity) {
	return func(e Entity) {
		e.assignComponents(components...)
	}
}

func (e *baseEntity) assignComponents(components ...Component) {
	e.components = components
}

func (e *baseEntity) ID() string {
	return e.id
}

func (e *baseEntity) Game() *Game {
	return e.game
}

func (e *baseEntity) AddComponent(c Component) error {
	if e.initialized {
		err := c.Init(e)
		if err != nil {
			return err
		}
	}

	e.mutex.Lock()
	e.components = append(e.components, c)
	e.mutex.Unlock()

	return nil
}

func (e *baseEntity) GetComponent(componentUID ksuid.KSUID) Component {
	for _, c := range e.components {
		if c.GetComponentUID() == componentUID {
			return c
		}
	}

	return nil
}

func (e *baseEntity) GetComponents(componentUID ksuid.KSUID) []Component {
	components := make([]Component, 0)

	for _, c := range e.components {
		if c.GetComponentUID() == componentUID {
			components = append(components, c)
		}
	}

	return components
}

func (e *baseEntity) RemoveComponent(component Component) bool {
	for i, c := range e.components {
		if c == component {
			e.components = append(e.components[:i], e.components[i+1:]...)
			return true
		}
	}

	return false
}

func (e *baseEntity) Clone() Entity {
	return NewBaseEntity(e.game, NewBaseEntityComponents(e.components...))
}

func (e *baseEntity) Init() error {
	for _, c := range e.components {
		if err := c.Init(e); err != nil {
			return err
		}
	}
	return nil
}

func (e *baseEntity) Dispose() {
	for _, c := range e.components {
		c.Dispose()
	}
}

func (e *baseEntity) Update() error {
	for _, c := range e.components {
		if err := c.PreUpdate(); err != nil {
			return err
		}
	}

	for _, c := range e.components {
		if err := c.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (e *baseEntity) Draw(screen *ebiten.Image) {
	for _, c := range e.components {
		c.DrawOn(screen)
	}
}
