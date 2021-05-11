package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/internal/game"
)

type System interface {
	Init(g game.Game) error
	Update(g game.Game) error
	Draw(g game.Game, screen *ebiten.Image)
}

type MergedSystem struct {
	Children []System
}

type system struct {
	init   func(g game.Game) error
	update func(g game.Game) error
	draw   func(g game.Game, screen *ebiten.Image)
}

func (b system) Init(g game.Game) error {
	if b.init == nil {
		return nil
	}

	return b.init(g)
}

func (b system) Update(g game.Game) error {
	if b.update == nil {
		return nil
	}

	return b.update(g)
}

func (b system) Draw(g game.Game, screen *ebiten.Image) {
	if b.draw == nil {
		return
	}

	b.draw(g, screen)
}

type SystemOptions struct {
	Init   func(g game.Game) error
	Update func(g game.Game) error
	Draw   func(g game.Game, screen *ebiten.Image)
}

func NewSystem(options *SystemOptions) System {
	return &system{
		init:   options.Init,
		update: options.Update,
		draw:   options.Draw,
	}
}

func MergeSystems(systems ...System) System {
	return NewSystem(&SystemOptions{
		Init: func(g game.Game) error {
			for _, s := range systems {
				if err := s.Init(g); err != nil {
					return err
				}
			}
			return nil
		},
		Update: func(g game.Game) error {
			for _, s := range systems {
				if err := s.Update(g); err != nil {
					return err
				}
			}
			return nil
		},
		Draw: func(g game.Game, screen *ebiten.Image) {
			for _, s := range systems {
				s.Draw(g, screen)
			}
		},
	})
}
