package component

import (
	"github.com/cstria0106/tamagotchi/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Drawable struct {
	X, Y, W, H int16
	Image      *ebiten.Image
	Color      color.Color
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
			Color: color.Black,
		},
	)
}
