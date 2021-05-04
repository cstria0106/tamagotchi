package drawable

import (
	"github.com/gopackage/tween/curves"
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/resources/images"
	"time"
)

type Character struct {
	image      *ebiten.Image
	x, y, w, h int16
}

func NewCharacter(x, y int16) *Character {
	image := ebiten.NewImageFromImage(images.Character)
	w, h := image.Size()
	character := Character{image: image, x: x, y: y, w: int16(w), h: int16(h)}
	return &character
}

func (c *Character) DrawOn(image *ebiten.Image) {
	var options ebiten.DrawImageOptions

	options.GeoM.Translate(float64(c.x-c.w/2), float64(c.y-c.h/2))

	image.DrawImage(c.image, &options)
}

func (c *Character) MoveTo(x, y int16) {
	originalX := c.x
	originalY := c.y

	go func() {
		step := 0.0
		for {
			c.x = originalX + int16(float64(x-originalX)*curves.EaseInOutSine(step/60))
			c.y = originalY + int16(float64(y-originalY)*curves.EaseInOutSine(step/60))

			time.Sleep(time.Second / 60)
			step += 1
			if step > 60 {
				break
			}
		}
		c.x = x
		c.y = y
	}()
}
