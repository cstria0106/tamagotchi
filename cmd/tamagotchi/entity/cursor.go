package entity

import (
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/component"
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/resources/images"
	"github.com/cstria0106/tamagotchi/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewCursor() *game.Entity {
	image := ebiten.NewImageFromImage(images.Cursor)
	w, h := image.Size()

	return game.NewEntity(game.NewEntityComponents(
		component.NewDrawable(0, 0, int16(w), int16(h), image),
		component.NewCursor(),
	))
}
