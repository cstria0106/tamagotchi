package system

import (
	"github.com/gopackage/tween/curves"
	"golang.org/x/exp/rand"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/tamagotchi/component"
)

func Character() *game.System {
	return &game.System{
		Init: func(g *game.Game) error {
			g.WithComponents([]game.ComponentID{component.CHARACTER, component.TWEEN}, func(components []*game.Component) {
				tween := components[1].Data.(*component.Tween)
				tween.Duration = 1000
				tween.Curve = curves.Linear
			})
			return nil
		},
		Update: func(g *game.Game) error {
			g.WithComponents(
				[]game.ComponentID{component.CHARACTER, component.DRAWABLE, component.MOUSEEVENT, component.TWEEN},
				func(components []*game.Component) {
					character := components[0].Data.(*component.Character)
					drawable := components[1].Data.(*component.Drawable)
					clickable := components[2].Data.(*component.MouseEvent)
					tween := components[3].Data.(*component.Tween)

					if clickable.JustDown {
						x := int16(rand.Uint64() % 100)
						y := int16(rand.Uint64() % 150)
						character.MoveTo(x, y)
					}

					if character.Moving {
						tween.Reset()
						tween.Resume()
						character.Moving = false
					}

					character.X = character.LastX + int16(float64(character.TargetX-character.LastX)*tween.Value)
					character.Y = character.LastY + int16(float64(character.TargetY-character.LastY)*tween.Value)

					drawable.X = character.X
					drawable.Y = character.Y
					clickable.MoveTo(character.X, character.Y)
				},
			)

			return nil
		},
	}
}
