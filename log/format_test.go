package log

import (
	"fmt"
	"testing"
)

func TestNewFormat(t *testing.T) {
	type row struct {
		given string
		exp   string
	}
	tt := []row{
		{
			given: "blah",
			exp:   "<nil>",
		},
		{
			given: "{TIME}{TAG}{MSG}",
			exp:   "*log.Format",
		},
	}
	for _, r := range tt {
		f := NewFormat(r.given)
		res := fmt.Sprintf("%T", f)
		if res != r.exp {
			t.Errorf("given %q expected %q got %q", r.given, r.exp, res)
		}
	}
}
func TestFormatDup(t *testing.T) {
	f1 := NewFormat("{TAG}{TIME}{MSG}")
	f2 := f1.Dup()
	if !compareFormat(f1, f2) {
		t.Errorf("given %v Dup expected to return %v but got %v", f1, f1, f2)
	}
}

func TestValidateFormat(t *testing.T) {
	type row struct {
		given string
		exp   bool
	}
	tt := []row{
		{
			given: "Foo Bar",
			exp:   false,
		},
		{
			given: "{TIME} {MSG}",
			exp:   false,
		},
		{
			given: "{MSG}",
			exp:   false,
		},
		{
			given: "{TIME} {TAG}",
			exp:   false,
		},
		{
			given: "{TAG}<MSG>{}{}",
			exp:   false,
		},
		{
			given: "{TAG}{TIME}{MSG}",
			exp:   true,
		},
		{
			given: "{TIME}{}}{{}{MSG}  {TAG}",
			exp:   true,
		},
	}
	for _, r := range tt {
		if res := validateFormat(r.given); res != r.exp {
			t.Errorf("given %q expected %v got %v", r.given, r.exp, res)
		}
	}
}

func compareFormat(f1, f2 IFormat) bool {
	return f1.Get() == f2.Get()
}
