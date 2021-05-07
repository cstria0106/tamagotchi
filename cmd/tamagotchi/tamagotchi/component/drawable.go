package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/segmentio/ksuid"
	"tamagotchi/cmd/tamagotchi/game"
)

var DrawableUid = ksuid.New()

type Drawable struct {
	game.BaseComponent
	X, Y, W, H int16
	Image      *ebiten.Image
}

func (d *Drawable) GetComponentUID() ksuid.KSUID {
	return DrawableUid
}

func (d *Drawable) Init(_ game.Entity) error {
	return nil
}

func (d *Drawable) DrawOn(screen *ebiten.Image) {
	var options ebiten.DrawImageOptions

	options.GeoM.Translate(float64(d.X-d.W/2), float64(d.Y-d.H/2))

	screen.DrawImage(d.Image, &options)
}

func (d *Drawable) Clone() game.Component {
	return &Drawable{X: d.X, Y: d.Y, W: d.W, H: d.H, Image: ebiten.NewImageFromImage(d.Image)}
}
