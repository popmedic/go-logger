package log

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tags := []string{
		defaultInfoTag,
		defaultDebugTag,
		defaultWarnTag,
		defaultErrorTag,
		defaultFatalTag,
	}
	if res := len(tiers); res != 5 {
		t.Errorf("tiers should %d is %d", 5, res)
	} else {
		for i, tag := range tags {
			if res := tiers[i].GetTag().Get(); res != tag {
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

	exp := fmt.Sprintf("<%s>: Test %s number %d", defaultInfoTag, defaultInfoTag, infoIdx)
	Info("Test", defaultInfoTag, "number", infoIdx)
	b := w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}
	w.Reset()
	Infof("Test %s number %d", defaultInfoTag, infoIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}

	w.Reset()
	exp = fmt.Sprintf("<%s>: Test %s number %d", defaultDebugTag, defaultDebugTag, debugIdx)
	Debug("Test", defaultDebugTag, "number", debugIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}
	w.Reset()
	Debugf("Test %s number %d", defaultDebugTag, debugIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}

	w.Reset()
	exp = fmt.Sprintf("<%s>: Test %s number %d", defaultWarnTag, defaultWarnTag, warnIdx)
	Warn("Test", defaultWarnTag, "number", warnIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}
	w.Reset()
	Warnf("Test %s number %d", defaultWarnTag, warnIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}

	w.Reset()
	exp = fmt.Sprintf("<%s>: Test %s number %d", defaultErrorTag, defaultErrorTag, errorIdx)
	Error("Test", defaultErrorTag, "number", errorIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}
	w.Reset()
	Errorf("Test %s number %d", defaultErrorTag, errorIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	}

	w.Reset()
	exitCalled := false
	exp = fmt.Sprintf("<%s>: Test %s number %d", defaultFatalTag, defaultFatalTag, fatalIdx)
	Fatal(func(int) { exitCalled = true }, "Test", defaultFatalTag, "number", fatalIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	} else if !exitCalled {
		t.Error("exit was not called")
	}
	w.Reset()
	exitCalled = false
	Fatalf(func(int) { exitCalled = true }, "Test %s number %d", defaultFatalTag, fatalIdx)
	b = w.Bytes()
	if !strings.Contains(string(b), exp) {
		t.Errorf("expected suffix %s got %s", exp, string(b))
	} else if !exitCalled {
		t.Error("exit was not called")
	}
}

func TestSetFormatFail(t *testing.T) {
	if err := SetFormat(""); err == nil {
		t.Errorf("given a bad format string expected error %q", invaledTimeFormatFmt)
	}
}
