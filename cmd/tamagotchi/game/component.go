package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/segmentio/ksuid"
)

type Component interface {
	GetComponentUID() ksuid.KSUID
	Clone() Component
	Init(e Entity) error
	PreUpdate() error
	Update() error
	DrawOn(_ *ebiten.Image)
	Dispose()
}

type BaseComponent struct{}

func (c *BaseComponent) GetComponentUID() ksuid.KSUID {
	panic(errors.New("get component uid is not implemented"))
}

func (c *BaseComponent) Clone() *Component {
	panic(errors.New("clone is not implemented"))
}

func (c *BaseComponent) Init(_ Entity) error    { return nil }
func (c *BaseComponent) PreUpdate() error       { return nil }
func (c *BaseComponent) Update() error          { return nil }
func (c *BaseComponent) DrawOn(_ *ebiten.Image) {}
func (c *BaseComponent) Dispose()               {}
