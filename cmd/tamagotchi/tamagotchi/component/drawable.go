package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/cmd/tamagotchi/game"
)

type Drawable struct {
	X, Y, W, H int16
	Image      *ebiten.Image
}

func NewDrawable(x, y, w, h int16, image *ebiten.Image) *game.Component {
	return game.NewComponent(
		DRAWABLE,
		&Drawable{
			X:     x,
			Y:     y,
			W:     w,
			H:     h,
			Image: image,
		},
	)
}
