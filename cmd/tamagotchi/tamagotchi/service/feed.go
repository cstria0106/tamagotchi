package service

import (
	"tamagotchi/cmd/tamagotchi/game"
)

type FeedService struct {
	game.Entity
}

func (s *FeedService) Init() error {
	return nil
}

func NewFeedService(g *game.Game) game.Entity {
	return &FeedService{
		Entity: game.NewBaseEntity(g),
	}
}
