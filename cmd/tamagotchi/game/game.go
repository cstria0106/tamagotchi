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

	mutex   sync.Mutex
	started bool

	entities *[256][]Entity
}

func (g *Game) Start() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.started {
		return errors.New("game is already started")
	}

	g.started = true

	err := ebiten.RunGame(g)
	return err
}

func (g *Game) Update() error {
	for _, entities := range g.entities {
		for _, e := range entities {
			err := e.Update()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, entities := range g.entities {
		for _, e := range entities {
			e.Draw(screen)
		}
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return int(g.screenOptions.Width), int(g.screenOptions.Height)
}

func (g *Game) AddOrderedEntity(order int16, e Entity) error {
	if order == 255 {
		return errors.New("order must be 0 ~ 254")
	}

	g.mutex.Lock()
	g.entities[1+order] = append(g.entities[1+order], e)
	g.mutex.Unlock()

	return e.Init()
}

func (g *Game) AddOrderedEntities(order int16, entities ...Entity) error {
	if order > 254 {
		return errors.New("order must be 0 ~ 254")
	}

	g.mutex.Lock()
	g.entities[1+order] = append(g.entities[1+order], entities...)
	g.mutex.Unlock()

	for _, e := range entities {
		err := e.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) AddEntity(e Entity) error {
	return g.AddOrderedEntity(0, e)
}

func (g *Game) AddEntities(entities ...Entity) error {
	return g.AddOrderedEntities(0, entities...)
}

func (g *Game) RemoveEntity(e Entity) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	for _, entities := range g.entities {
		for i, element := range entities {
			if element == e {
				element.Dispose()
				entities = append(entities[:i], entities[i+1:]...)
				return true
			}
		}
	}

	return false
}

func (g *Game) AddService(s Entity) error {
	return g.AddOrderedEntity(-1, s)
}

func (g *Game) ApplyScreenOptions(options *ScreenOptions) {
	g.screenOptions = options
	ebiten.SetWindowSize(int(options.Width*options.Scale), int(options.Height*options.Scale))
}

func New(o *Options) (*Game, error) {
	rand.Seed(uint64(time.Now().Unix()))

	entities := [256][]Entity{}

	// Services (1024)
	entities[0] = make([]Entity, 0, 1024)

	// Most bottom layer entities (1024)
	entities[1] = make([]Entity, 0, 1024)

	// Other entities (128)
	for i := 2; i < 256; i++ {
		entities[i] = make([]Entity, 0, 128)
	}

	g := &Game{
		screenOptions: o.Screen,
		client:        o.Client,

		mutex:   sync.Mutex{},
		started: false,

		entities: &entities,
	}

	g.ApplyScreenOptions(o.Screen)

	return g, nil
}
