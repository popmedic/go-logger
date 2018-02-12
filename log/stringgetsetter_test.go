package log

import "testing"

func TestStringGetSetterGetSet(t *testing.T) {
	exp := "string"
	s := &StringGetSetter{
		str: "string2",
	}
	s.Set(exp)
	if res := s.Get(); res != exp {
		t.Errorf("given %q expected Set/Get to be %q got %q", exp, exp, res)
	}
}
