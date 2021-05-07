package component

import (
	"github.com/gopackage/tween/curves"
	"github.com/segmentio/ksuid"
	"tamagotchi/cmd/tamagotchi/game"
)

var TweenUID = ksuid.New()

type TweenOptions struct {
	Loop       bool
	Duration   uint64
	OnComplete func()
	OnChange   func(float64)
	Curve      func(float64) float64
}

type Tween struct {
	game.BaseComponent

	playing bool
	tick    uint64

	loop       bool
	duration   uint64
	onComplete func()
	onChange   func(float64)
	curve      func(float64) float64

	entity game.Entity
}

func (t *Tween) Init(e game.Entity) error {
	t.entity = e
	return nil
}

func (t *Tween) OnComplete(f func()) {
	t.onComplete = f
}

func (t *Tween) Loop(loop bool) {
	t.loop = loop
}

func (t *Tween) SetProgress(progress uint16) {
	t.tick = uint64(progress) * (t.duration * 60 / 1000)
}

func (t *Tween) Reset() {
	t.SetProgress(0)
	t.Pause()
}

func (t *Tween) Start() {
	t.Reset()
	t.Resume()
}

func (t *Tween) Pause() {
	t.playing = false
}

func (t *Tween) Resume() {
	t.playing = true
}

func (t *Tween) GetComponentUID() ksuid.KSUID {
	return TweenUID
}

func (t *Tween) Clone() game.Component {
	return &Tween{
		BaseComponent: t.BaseComponent,
		entity:        t.entity,
		playing:       t.playing,
		tick:          t.tick,
		loop:          t.loop,
		duration:      t.duration,
		onChange:      t.onChange,
		onComplete:    t.onComplete,
		curve:         t.curve,
	}
}

func (t *Tween) PreUpdate() error {
	progress := float64(t.tick) / (float64(t.duration) * 60 / 1000)

	if t.playing {
		t.onChange(t.curve(progress))
		t.tick += 1
	}

	if progress >= 1 {
		t.onChange(1)

		t.entity.RemoveComponent(t)

		if t.loop {
			newTween := t.Clone().(*Tween)

			newTween.Reset()
			newTween.Start()
			_ = t.entity.AddComponent(newTween)
		}
	}

	return nil
}

func NewTween(options *TweenOptions) *Tween {
	curve := options.Curve
	if curve == nil {
		curve = curves.Linear
	}

	return &Tween{
		loop:       options.Loop,
		duration:   options.Duration,
		onChange:   options.OnChange,
		onComplete: options.OnComplete,
		curve:      options.Curve,
	}
}
