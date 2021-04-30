package game

import (
	"github.com/Tarliton/collision2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
)

type Button struct {
	image      *ebiten.Image
	polygon    collision2d.Polygon
	x, y, w, h float64
	Down bool
	JustDown bool
	JustUp bool
}

func NewButton(image *ebiten.Image, x, y, w, h float64) *Button {
	imageSizeW, imageSizeH := image.Size()

	collider := collision2d.NewBox(collision2d.Vector{X: x, Y: y}, float64(imageSizeW) * w, float64(imageSizeH) * h)

	button := Button{
		image:   image,
		polygon: collider.ToPolygon(),
		x:       x,
		y:       y,
		w:       w,
		h:       h,
	}

	return &button
}

func NewButtonFromImage(image image.Image, x,y,w,h float64) *Button  {
	eImage := ebiten.NewImageFromImage(image)
	return NewButton(eImage, x, y, w, h)
}

func (b *Button) CheckClick() {
	b.JustDown = false
	b.JustUp = false

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if collision2d.PointInPolygon(collision2d.Vector{X: float64(x), Y: float64(y)}, b.polygon) {
			b.Down = true

			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				b.JustDown = true
			}
		}

		return
	}


	if b.Down {
		b.JustUp = true
		b.Down = false
	}
}

func (b *Button) DrawOn(image *ebiten.Image) {
	var options ebiten.DrawImageOptions

	options.GeoM.Scale(b.w, b.h)
	options.GeoM.Translate(b.x, b.y)

	image.DrawImage(b.image, &options)
}
