package log

import (
	"testing"
)

func TestNewTag(t *testing.T) {
	exp := "TAG"
	tg := NewTag(exp)
	if nil == t {
		t.Error("given a call to NewTag expected not nil got nil")
	} else if res := tg.Get(); res != exp {
		t.Errorf("given a call to NewTag with %q expected a new tag with a tag %q got %v", exp, exp, res)
	}
}

func TestTagDup(t *testing.T) {
	exp := NewTag("TAG")
	if res := exp.Dup(); !compareTag(exp, res) {
		t.Errorf("given %v expected %v got %v", exp, exp, res)
	}
}

func compareTag(t1, t2 ITag) bool {
	return t1.Get() == t2.Get()
}
