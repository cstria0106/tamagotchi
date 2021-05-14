package component

import (
	"github.com/cstria0106/tamagotchi/internal/game"
	"github.com/gopackage/tween/curves"
)

type Tween struct {
	Playing bool
	Tick    uint64
	Value   float64

	Loop     bool
	Duration uint64
	Curve    func(float64) float64
}

type TweenOptions struct {
	Loop     bool
	Duration uint64
	Curve    func(float64) float64
}

func (t *Tween) Pause() {
	t.Playing = false
}

func (t *Tween) Resume() {
	t.Playing = true
}

func (t *Tween) Reset() {
	t.Tick = 0
	t.Value = t.Curve(0)
}

func NewTween(options *TweenOptions) *game.Component {
	var loop bool
	var duration uint64
	var curve func(float64) float64

	if options != nil {
		loop = options.Loop
		duration = options.Duration
		curve = options.Curve
	}

	if curve == nil {
		curve = curves.Linear
	}

	return game.NewComponent(
		TWEEN,
		&Tween{
			Loop:     loop,
			Duration: duration,
			Curve:    curve,
		},
	)
}
