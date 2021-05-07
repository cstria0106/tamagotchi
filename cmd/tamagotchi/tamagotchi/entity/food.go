package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/resources/images"
	"tamagotchi/cmd/tamagotchi/tamagotchi/component"
)

type Food struct {
	game.Entity
}

func NewFood(g *game.Game, x, y int16) game.Entity {
	image := ebiten.NewImageFromImage(images.Food)
	w, h := image.Size()

	return Food{
		Entity: g.NewBaseEntity(game.NewBaseEntityComponents(
			&component.Drawable{
				X: x, Y: y, W: int16(w), H: int16(h), Image: image,
			},
		)),
	}
}
