package log

import (
	"sync"
)

// Color is the type used for all colors
type Color struct {
	StringGetSetter
	end StringGetSetter
}

// NewColor creates a new Color
func NewColor(color, endColor string) IColor {
	return &Color{
		StringGetSetter{
			str:  color,
			lock: sync.RWMutex{},
		},
		StringGetSetter{
			str:  endColor,
			lock: sync.RWMutex{},
		},
	}
}

// End gets the suffix string for the color
func (c *Color) End() string {
	return c.end.Get()
}

// Dup duplicates the Color
func (c *Color) Dup() IColor {
	return NewColor(c.Get(), c.End())
}
