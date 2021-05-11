package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"tamagotchi/cmd/tamagotchi/engine"
	"tamagotchi/internal/game"
)

var FPSCounter = engine.NewSystem(&engine.SystemOptions{
	Draw: func(_ game.Game, screen *ebiten.Image) {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%.0f", ebiten.CurrentFPS()))
	},
})
