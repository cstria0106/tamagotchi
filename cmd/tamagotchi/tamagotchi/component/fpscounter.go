package component

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/segmentio/ksuid"
	"tamagotchi/cmd/tamagotchi/game"
)

var FpsCounterUid = ksuid.New()

type FpsCounter struct {
	game.BaseComponent
}

func (c *FpsCounter) GetComponentUID() ksuid.KSUID {
	return FpsCounterUid
}

func (c *FpsCounter) DrawOn(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%.0f", ebiten.CurrentFPS()))
}

func (c FpsCounter) Clone() game.Component {
	return &FpsCounter{}
}
