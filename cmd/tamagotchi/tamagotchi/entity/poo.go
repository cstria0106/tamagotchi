package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/resources/images"
	"tamagotchi/cmd/tamagotchi/tamagotchi/component"
)

type Poo struct {
	game.Entity
}

func NewPoo(g *game.Game, x, y int16) game.Entity {
	image := ebiten.NewImageFromImage(images.Food)
	w, h := image.Size()

	return Poo{
		Entity: g.NewBaseEntity(game.NewBaseEntityComponents(
			&component.Drawable{
				X: x, Y: y, W: int16(w), H: int16(h), Image: image,
			},
		)),
	}
}
