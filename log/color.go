package log

import (
	"sync"

	"github.com/popmedic/go-color/colorize"
)

type Color struct {
	colorize.IColorize
	lock sync.RWMutex
}

func NewColor(cz colorize.IColorize) IColor {
	return &Color{
		cz,
		sync.RWMutex{},
	}
}

func (c *Color) Get() IColor {
	return c.Dup()
}

func (c *Color) Set(color IColor) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.IColorize = color
}

func (c *Color) Dup() IColor {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return NewColor(c)
}
