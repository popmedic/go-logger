package log

import "testing"

func TestNewTimeFormat(t *testing.T) {
	exp := "blah"
	f := NewTimeFormat(exp)
	if f == nil {
		t.Errorf("should have created a new time format, got %v", f)
	} else if format := f.Get(); format != exp {
		t.Errorf("given a new time format of %q expected a new time format with Get value %q got %q", exp, exp, format)
	}
}
func TestTimeFormatDup(t *testing.T) {
	f1 := NewTimeFormat("blah")
	f2 := f1.Dup()
	if !compareTimeFormat(f1, f2) {
		t.Errorf("given %+v Dup expected to return %+v but got %+v", f1, f1, f2)
	}
}

func compareTimeFormat(f1, f2 ITimeFormat) bool {
	return f1.Get() == f2.Get()
}
