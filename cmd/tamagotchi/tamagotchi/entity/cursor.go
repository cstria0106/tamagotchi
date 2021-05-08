package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/resources/images"
	"tamagotchi/cmd/tamagotchi/tamagotchi/component"
)

func NewCursor() *game.Entity {
	image := ebiten.NewImageFromImage(images.Cursor)
	w, h := image.Size()

	return game.NewEntity(game.NewEntityComponents(
		component.NewDrawable(0, 0, int16(w), int16(h), image),
		component.NewCursor(),
	))
}
