package entity

import (
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/component"
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/resources/images"
	"github.com/cstria0106/tamagotchi/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewCharacter(x, y int16) *game.Entity {
	image := ebiten.NewImageFromImage(images.Character)
	w, h := image.Size()

	return game.NewEntity(game.NewEntityComponents(
		component.NewCharacter(x, y),
		component.NewDrawable(x, y, int16(w), int16(h), image),
		component.NewClickable(x, y, int16(w), int16(h)),
		component.NewTween(nil),
	))
}
