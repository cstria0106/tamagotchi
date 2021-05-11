package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/cmd/tamagotchi/component"
	"tamagotchi/cmd/tamagotchi/resources/images"
	"tamagotchi/internal/game"
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
