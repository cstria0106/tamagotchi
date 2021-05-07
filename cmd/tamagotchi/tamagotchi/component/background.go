package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/segmentio/ksuid"
	"image/color"
	"tamagotchi/cmd/tamagotchi/game"
)

var BackgroundUid = ksuid.New()

type Background struct {
	game.BaseComponent
	Color color.Color
}

func (b *Background) GetComponentUID() ksuid.KSUID {
	return BackgroundUid
}

func (b *Background) DrawOn(screen *ebiten.Image) {
	screen.Fill(b.Color)
}

func (b *Background) Clone() game.Component {
	return &Background{Color: b.Color}
}
