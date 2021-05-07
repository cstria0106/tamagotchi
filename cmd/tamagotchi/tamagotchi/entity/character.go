package entity

import (
	"github.com/gopackage/tween/curves"
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/resources/images"
	"tamagotchi/cmd/tamagotchi/tamagotchi/component"
)

type Character struct {
	game.Entity
}

func NewCharacter(g *game.Game, x, y int16) *Character {
	image := ebiten.NewImageFromImage(images.Character)
	w, h := image.Size()

	return &Character{
		game.NewBaseEntity(g, game.NewBaseEntityComponents(
			&component.Drawable{X: x, Y: y, W: int16(w), H: int16(h), Image: image},
		)),
	}
}

func (c *Character) MoveTo(x, y int16) {
	drawable := c.GetComponent(component.DrawableUid).(*component.Drawable)
	initialX := drawable.X
	initialY := drawable.Y

	tween := component.NewTween(&component.TweenOptions{
		OnChange: func(u float64) {
			drawable.X = initialX + int16(float64(x-initialX)*u)
			drawable.Y = initialY + int16(float64(y-initialY)*u)
		},
		Duration: 200,
		Curve:    curves.Linear,
		Loop:     true,
	})

	_ = c.AddComponent(tween)

	tween.Start()
}
