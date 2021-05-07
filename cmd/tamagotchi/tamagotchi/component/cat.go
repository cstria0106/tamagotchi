package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/segmentio/ksuid"
	"tamagotchi/cmd/tamagotchi/game"
)

var CatUID = ksuid.New()

type Cat struct {
	game.BaseComponent
}

func (c *Cat) GetComponentUID() ksuid.KSUID {
	return CatUID
}

func (c *Cat) Init(_ game.Entity) error {
	return nil
}

func (c *Cat) DrawOn(_ *ebiten.Image) {

}

func (c *Cat) Clone() game.Component {
	return &Cat{}
}
