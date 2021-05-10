package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"tamagotchi/cmd/tamagotchi/game"
)

func FPSCounter() *game.System {
	return &game.System{
		Draw: func(_ *game.Game, screen *ebiten.Image) {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("%.0f", ebiten.CurrentFPS()))
		},
	}
}
