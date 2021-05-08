package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/tamagotchi/component"
)

func Character() *game.System {
	return &game.System{
		ID: CHARACTER,
		Init: func(g *game.Game) error {
			return nil
		},
		Update: func(g *game.Game) error {
			g.WithComponents(
				[]game.ComponentID{component.CHARACTER, component.DRAWABLE, component.MOUSEEVENT},
				func(components []*game.Component) {
					character := components[0].Data.(*component.Character)
					drawable := components[1].Data.(*component.Drawable)
					clickable := components[2].Data.(*component.MouseEvent)

					cursorX, cursorY := ebiten.CursorPosition()
					centerX := drawable.X + drawable.W/2
					centerY := drawable.Y + drawable.H/2

					deltaX := centerX - int16(cursorX)
					deltaY := centerY - int16(cursorY)

					direction := math.Atan2(float64(deltaY), float64(deltaX))
					distance := math.Sqrt(float64(deltaX*deltaX + deltaY*deltaY))

					character.X += math.Max(1, 5-(distance/10)) * math.Cos(direction)
					character.Y += math.Max(1, 5-(distance/10)) * math.Sin(direction)

					if character.X < 0 {
						character.X = 0
					}

					if character.X > float64(100-drawable.W) {
						character.X = float64(100 - drawable.W)
					}

					if character.Y < 0 {
						character.Y = 0
					}

					if character.Y > float64(150-drawable.H) {
						character.Y = float64(150 - drawable.H)
					}

					drawable.X = int16(character.X)
					drawable.Y = int16(character.Y)
					clickable.MoveTo(int16(character.X), int16(character.Y))
				},
			)

			return nil
		},
	}
}
