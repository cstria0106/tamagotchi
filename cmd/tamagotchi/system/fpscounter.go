package system

import (
	"fmt"
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/engine"
	"github.com/cstria0106/tamagotchi/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var FPSCounter = engine.NewSystem(&engine.SystemOptions{
	Draw: func(_ game.Game, screen *ebiten.Image) {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%.0f", ebiten.CurrentFPS()))
	},
})
