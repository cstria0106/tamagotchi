package drawable

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/resources/images"
)

type Poo struct {
	ID         uint32
	image      *ebiten.Image
	x, y, w, h int16
}

func NewPoo(id uint32, x, y int16) *Poo {
	image := ebiten.NewImageFromImage(images.Poo)
	w, h := image.Size()
	poo := Poo{ID: id, image: image, x: x, y: y, w: int16(w), h: int16(h)}
	return &poo
}

func (c *Poo) DrawOn(image *ebiten.Image) {
	var options ebiten.DrawImageOptions

	options.GeoM.Translate(float64(c.x-c.w/2), float64(c.y-c.h/2))

	image.DrawImage(c.image, &options)
}

func (c *Poo) CheckHover() {

}
