package game

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	DrawOn(image *ebiten.Image)
}
