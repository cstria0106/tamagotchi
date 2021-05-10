package game

import (
	"fmt"
)

type Entity struct {
	components []*Component
}

func (g *Game) NewEntity(options ...func(*Entity)) *Entity {
	return NewEntity(options...)
}

func NewEntity(options ...func(*Entity)) *Entity {
	entity := Entity{
		components: make([]*Component, 0),
	}

	for _, option := range options {
		option(&entity)
	}

	return &entity
}

func NewEntityComponents(components ...*Component) func(*Entity) {
	return func(e *Entity) {
		e.AddComponents(components...)
	}
}

func (e *Entity) AddComponents(c ...*Component) {
	e.components = append(e.components, c...)
}

func (e *Entity) Components(componentIDs []ComponentID) (components []*Component, all bool) {
	count := 0
	length := len(componentIDs)

	components = make([]*Component, length)

	for _, c := range e.components {
		for index, id := range componentIDs {
			if c.ID == id {
				components[index] = c
				count++

				if count == length {
					return components, true
				}
			}
		}
	}

	return components, false
}

func (e *Entity) RemoveComponent(component *Component) (removed bool) {
	for i, c := range e.components {
		if c == component {
			e.components = append(e.components[:i], e.components[i+1:]...)
			return true
		}
	}

	return false
}

func (e *Entity) Clone() *Entity {
	components := make([]*Component, len(e.components))
	for i, c := range components {
		cloned := *c
		components[i] = &cloned
	}
	return NewEntity(NewEntityComponents(components...))
}

func (e *Entity) String() string {
	return fmt.Sprintf("entity(+%d components)", len(e.components))
}
