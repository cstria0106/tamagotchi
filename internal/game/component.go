package game

import (
	"strconv"
)

type ComponentID uint64

type Component struct {
	ID   ComponentID
	Data interface{}
}

func NewComponent(id ComponentID, data interface{}) *Component {
	return &Component{id, data}
}

func (c Component) String() string {
	return strconv.FormatUint(uint64(c.ID), 10)
}
