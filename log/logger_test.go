package log

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tags := []string{
		defaultInfoTag.Get(),
		defaultDebugTag.Get(),
		defaultWarnTag.Get(),
		defaultErrorTag.Get(),
		defaultFatalTag.Get(),
	}
	if res := len(tiers.tiers); res != 5 {
		t.Errorf("tiers should %d is %d", 5, res)
	} else {
		for i, tag := range tags {
			if res := tiers.Get(i).GetTag().Get(); res != tag {
				t.Errorf("tier %d should have tag %q has %q", i, tag, res)
			}
		}
	}
}

func TestLoggerOut(t *testing.T) {
	w := bytes.NewBuffer([]byte{})
	SetOutput(w)
	SetFormat("{TIME} <{TAG}>: {MSG}")
	SetTimeFormat("Mon Jan _2 15:04:05 2006")
	SetHTMLStatus(true)

	exp := fmt.Sprintf("<%s>: Test %s number %d", defaultInfoTag.Get(), defaultInfoTag.Get(), infoIdx)
	Info("Test", defaultInfoTag.Get(), "number", infoIdx)
	b := w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}
	w.Reset()
	Infof("Test %s number %d", defaultInfoTag.Get(), infoIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}

	w.Reset()
	exp = fmt.Sprintf("<%s>: Test %s number %d", defaultDebugTag.Get(), defaultDebugTag.Get(), debugIdx)
	Debug("Test", defaultDebugTag.Get(), "number", debugIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}
	w.Reset()
	Debugf("Test %s number %d", defaultDebugTag.Get(), debugIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}

	w.Reset()
	exp = fmt.Sprintf("<%s>: Test %s number %d", defaultWarnTag.Get(), defaultWarnTag.Get(), warnIdx)
	Warn("Test", defaultWarnTag.Get(), "number", warnIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}
	w.Reset()
	Warnf("Test %s number %d", defaultWarnTag.Get(), warnIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}

	w.Reset()
	exp = fmt.Sprintf("<%s>: Test %s number %d", defaultErrorTag.Get(), defaultErrorTag.Get(), errorIdx)
	Error("Test", defaultErrorTag.Get(), "number", errorIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}
	w.Reset()
	Errorf("Test %s number %d", defaultErrorTag.Get(), errorIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}

	w.Reset()
	exitCalled := false
	exp = fmt.Sprintf("<%s>: Test %s number %d", defaultFatalTag.Get(), defaultFatalTag.Get(), fatalIdx)
	Fatal(func(int) { exitCalled = true }, "Test", defaultFatalTag.Get(), "number", fatalIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	} else if !exitCalled {
		t.Error("exit was not called")
	}
	w.Reset()
	exitCalled = false
	Fatalf(func(int) { exitCalled = true }, "Test %s number %d", defaultFatalTag.Get(), fatalIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	} else if !exitCalled {
		t.Error("exit was not called")
	}

	expStr := `<!DOCTYPE html>
<html>
<body style="background-color:Black;">
<span style="color:White;"><span style="background-color:Black;"><span style="font-family:'Consola', monospace;">Fri Feb 16 20:21:50 2018 <INFO>: Test INFO number 0</span></span></span>
<br/>
<span style="color:White;"><span style="background-color:Black;"><span style="font-family:'Consola', monospace;">Fri Feb 16 20:21:50 2018 <INFO>: Test INFO number 0</span></span></span>
<br/>
<span style="color:Green;"><span style="background-color:Black;"><span style="font-family:'Consola', monospace;">Fri Feb 16 20:21:50 2018 <DEBUG>: Test DEBUG number 1</span></span></span>
<br/>
<span style="color:Green;"><span style="background-color:Black;"><span style="font-family:'Consola', monospace;">Fri Feb 16 20:21:50 2018 <DEBUG>: Test DEBUG number 1</span></span></span>
<br/>
<span style="color:Yellow;"><span style="background-color:Black;"><span style="font-family:'Consola', monospace;">Fri Feb 16 20:21:50 2018 <WARN>: Test WARN number 2</span></span></span>
<br/>
<span style="color:Yellow;"><span style="background-color:Black;"><span style="font-family:'Consola', monospace;">Fri Feb 16 20:21:50 2018 <WARN>: Test WARN number 2</span></span></span>
<br/>
<span style="color:Red;"><span style="background-color:Black;"><span style="font-family:'Consola', monospace;">Fri Feb 16 20:21:50 2018 <ERROR>: Test ERROR number 3</span></span></span>
<br/>
<span style="color:Red;"><span style="background-color:Black;"><span style="font-family:'Consola', monospace;">Fri Feb 16 20:21:50 2018 <ERROR>: Test ERROR number 3</span></span></span>
<br/>
<span style="background-color:Red;"><span style="color:White;"><span style="font-weight:bold;"><span style="font-family:'Consola', monospace;">Fri Feb 16 20:21:50 2018 <FATAL>: Test FATAL number 4</span></span></span></span>
<br/>
<span style="background-color:Red;"><span style="color:White;"><span style="font-weight:bold;"><span style="font-family:'Consola', monospace;">Fri Feb 16 20:21:50 2018 <FATAL>: Test FATAL number 4</span></span></span></span>
<br/>
</body>
</html>`
	if gotStr := htmlLoggerString(); strings.Compare(gotStr, expStr) == 0 {
		t.Errorf("expected:\n %s \ngot:\n %s", expStr, gotStr)
	}
}

type Tiers []ITier

func (tt Tiers) String() string {
	out := "["
	for _, t := range tt {
		out = fmt.Sprintf("%s %q", out, t.GetTag().Get())
	}
	return out + " ]"
}
