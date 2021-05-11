package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/cmd/tamagotchi/cli"
	"tamagotchi/internal/game"
)

type ScreenOptions struct {
	Width  int
	Height int
	Scale  int
}

type Options struct {
	Arguments *cli.Arguments
	Screen    *ScreenOptions
}

type Game struct {
	game.Game
	options *Options

	systems []System
}

func (g *Game) AddSystem(systems ...game.System) error {
	err := g.Game.AddSystem(systems...)
	if err != nil {
		return err
	}

	for _, s := range systems {
		s, ok := s.(System)

		if ok && s.Draw != nil {
			g.systems = append(g.systems, s)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, s := range g.systems {
		s.Draw(g, screen)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return g.options.Screen.Width, g.options.Screen.Height
}

func (g *Game) ApplyScreenOptions(options *ScreenOptions) {
	g.options.Screen = options
	ebiten.SetWindowSize(options.Width*options.Scale, options.Height*options.Scale)
}

func NewGame(options *Options) (*Game, error) {
	g, err := game.New()

	if err != nil {
		return nil, err
	}

	return &Game{
		Game:    g,
		options: options,
		systems: make([]System, 0, 1024),
	}, nil
}
