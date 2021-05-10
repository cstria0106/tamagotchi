package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/rand"
	"sync"
	"tamagotchi/cmd/tamagotchi/client"
	"time"
)

type ScreenOptions struct {
	Width  uint16
	Height uint16
	Scale  uint16
}

type Options struct {
	Client *client.Client
	Screen *ScreenOptions
}

type Game struct {
	screenOptions *ScreenOptions
	client        *client.Client

	lock    sync.Mutex
	started bool

	entities []*Entity
	systems  []*System
}

func (g *Game) Client() *client.Client {
	return g.client
}

func (g *Game) Start() error {
	g.lock.Lock()
	if g.started {
		return errors.New("game already started")
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

	err := ebiten.RunGame(g)
	return err
}

func (g *Game) Update() error {
	for _, system := range g.systems {
		if system.Update != nil {
			if err := system.Update(g); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, system := range g.systems {
		if system.Draw != nil {
			system.Draw(g, screen)
		}
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return int(g.screenOptions.Width), int(g.screenOptions.Height)
}

func (g *Game) WithComponents(componentIds []ComponentID, f func(components []*Component)) {
	for _, e := range g.entities {
		components, all := e.Components(componentIds)

		if all {
			f(components)
		}
	}
}

func (g *Game) AddEntities(entity ...*Entity) {
	g.entities = append(g.entities, entity...)
}

func (g *Game) RemoveEntity(entity *Entity) bool {
	for i, e := range g.entities {
		if entity == e {
			g.entities = append(g.entities[:i], g.entities[i+1:]...)
			return true
		}
	}

	return false
}

func (g *Game) AddSystem(systems ...*System) error {
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

func (g *Game) ApplyScreenOptions(options *ScreenOptions) {
	g.screenOptions = options
	ebiten.SetWindowSize(int(options.Width*options.Scale), int(options.Height*options.Scale))
}

func New(o *Options) (*Game, error) {
	rand.Seed(uint64(time.Now().Unix()))

	g := &Game{
		screenOptions: o.Screen,
		client:        o.Client,

		entities: make([]*Entity, 0, 1024),
		systems:  make([]*System, 0, 1024),
	}

	g.ApplyScreenOptions(o.Screen)

	return g, nil
}
