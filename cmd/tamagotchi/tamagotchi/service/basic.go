package service

import (
	"image"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/tamagotchi/component"
)

type BasicService struct {
	game.Entity
}

func (s *BasicService) Init() error {
	g := s.Game()

	_ = g.AddOrderedEntities(
		0,
		g.NewBaseEntity(game.NewBaseEntityComponents(
			&component.Background{Color: image.White},
		)),
	)

	_ = g.AddOrderedEntities(
		254,
		g.NewBaseEntity(game.NewBaseEntityComponents(
			&component.FpsCounter{},
		)),
	)

	return nil
}

func NewBasicService(g *game.Game) game.Entity {
	return &BasicService{
		Entity: game.NewBaseEntity(g),
	}
}
