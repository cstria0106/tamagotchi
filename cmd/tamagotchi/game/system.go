package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"strconv"
)

type SystemID uint64

const MERGED SystemID = 0

type System struct {
	ID     SystemID
	Game   *Game
	Init   func(g *Game) error
	Update func(g *Game) error
	Draw   func(g *Game, screen *ebiten.Image)
}

func (s *System) String() string {
	return strconv.FormatUint(uint64(s.ID), 10)
}

func MergeSystems(systems ...*System) *System {
	return &System{
		ID: MERGED,
		Init: func(g *Game) error {
			return g.AddSystem(systems...)
		},
	}
}
