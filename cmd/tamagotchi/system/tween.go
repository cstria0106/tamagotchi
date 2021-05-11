package system

import (
	"tamagotchi/cmd/tamagotchi/component"
	"tamagotchi/cmd/tamagotchi/engine"
	"tamagotchi/internal/game"
)

var Tween = engine.NewSystem(&engine.SystemOptions{
	Update: func(g game.Game) error {
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
})
