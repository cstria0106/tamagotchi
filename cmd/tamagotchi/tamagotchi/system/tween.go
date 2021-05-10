package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/tamagotchi/component"
)

func Tween() *game.System {
	return &game.System{
		Draw: func(_ *game.Game, screen *ebiten.Image) {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("%.0f", ebiten.CurrentFPS()))
		},
		Update: func(g *game.Game) error {
			g.WithComponents([]game.ComponentID{component.TWEEN}, func(components []*game.Component) {
				tween := components[0].Data.(*component.Tween)

				progress := float64(tween.Tick) / (float64(tween.Duration) * 60 / 1000)

				if tween.Playing {
					tween.Value = tween.Curve(progress)
					tween.Tick++
				}

				if progress >= 1 {
					tween.Value = 1
					tween.Tick = 0

					if !tween.Loop {
						tween.Playing = false
					}
				}
			})

			return nil
		},
	}
}
