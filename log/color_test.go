package log

import (
	"testing"
)

func TestNewColor(t *testing.T) {
	exp := "yellow"
	expEnd := "none"
	c := NewColor(exp, expEnd)
	if c == nil {
		t.Errorf("should have created a new color object, got %v", c)
	} else if color := c.Get(); color != exp {
		t.Errorf("given a new color of %q expected a new color with Get value %q got %q", exp, exp, color)
	} else if color := c.End(); color != expEnd {
		t.Errorf("given a new color of %q expected a new color with Get value %q got %q", expEnd, expEnd, color)
	}
}
func TestColorDup(t *testing.T) {
	c1 := NewColor("yellow", "none")
	c2 := c1.Dup()
	if !compareColor(c1, c2) {
		t.Errorf("given %v Dup expected to return %v but got %v", c1, c1, c2)
	}
}

func compareColor(c1, c2 IColor) bool {
	return c1.Get() == c2.Get() &&
		c1.End() == c2.End()
}
