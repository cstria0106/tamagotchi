package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/resources/images"
)

type Character struct {
	image      *ebiten.Image
	x, y, w, h float64
}

func NewCharacter(x, y float64) *Character {
	image := ebiten.NewImageFromImage(images.Character)
	w, h := image.Size()
	character := Character{image: image, x: x, y: y, w: float64(w), h: float64(h)}
	return &character
}

func (c *Character) DrawOn(image *ebiten.Image) {
	var options ebiten.DrawImageOptions

	options.GeoM.Translate(c.x-c.w/2, c.y-c.h/2)

	image.DrawImage(c.image, &options)
}
