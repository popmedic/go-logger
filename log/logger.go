// Package log provides a multi-tiered logger used to log output to a io.Writer.
// The different tiers are (least to highest piority):
//  1. Info
//  2. Debug
//  3. Warn
//  4. Error
//  5. Fatal
// One can set the output, output format, output time format, tags, and colors for tags.
// One can also add more tiers, and create their own logger with all the different tiers wanted.
package log

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/popmedic/go-color/colorize/tty"
)

const invaledTimeFormatFmt = "%q is does not contain needed key words \"{TAG}\", \"{TIME}\", and \"{MSG}\""

const (
	infoIdx int = iota
	debugIdx
	warnIdx
	errorIdx
	fatalIdx
)

const (
	defaultFormatString     = "{TIME} [{TAG}]> {MSG}"
	defaultTimeFormatString = "2006-01-02 15:04:05"
)

const (
	defaultInfoTagString  = "INFO"
	defaultDebugTagString = "DEBUG"
	defaultWarnTagString  = "WARN"
	defaultErrorTagString = "ERROR"
	defaultFatalTagString = "FATAL"
)

var (
	defaultOut = os.Stdout

	defaultFormat     = NewFormat(defaultFormatString)
	defaultTimeFormat = NewTimeFormat(defaultTimeFormatString)
)

var (
	defaultInfoColor  = tty.Reset()
	defaultDebugColor = tty.FgGreen()
	defaultWarnColor  = tty.FgYellow()
	defaultErrorColor = tty.FgRed()
	defaultFatalColor = tty.BgRed().Add(tty.FgHiWhite())
)

var (
	defaultInfoTag  = NewTag(defaultInfoTagString)
	defaultDebugTag = NewTag(defaultDebugTagString)
	defaultWarnTag  = NewTag(defaultWarnTagString)
	defaultErrorTag = NewTag(defaultErrorTagString)
	defaultFatalTag = NewTag(defaultFatalTagString)
)

var (
	infoTier = NewTier(
		defaultInfoColor,
		defaultInfoTag,
		defaultFormat,
		defaultTimeFormat,
		defaultOut,
	)
	debugTier = NewTier(
		defaultDebugColor,
		defaultDebugTag,
		defaultFormat,
		defaultTimeFormat,
		defaultOut,
	)
	warnTier = NewTier(
		defaultWarnColor,
		defaultWarnTag,
		defaultFormat,
		defaultTimeFormat,
		defaultOut,
	)
	errorTier = NewTier(
		defaultErrorColor,
		defaultErrorTag,
		defaultFormat,
		defaultTimeFormat,
		defaultOut,
	)
	fatalTier = NewTier(
		defaultFatalColor,
		defaultFatalTag,
		defaultFormat,
		defaultTimeFormat,
		defaultOut,
	)
)

var (
	tiers = []ITier{
		infoTier,
		debugTier,
		warnTier,
		errorTier,
		fatalTier,
	}

	lock = sync.RWMutex{}
)

// SetOutput sets where the logger will write to
func SetOutput(out io.Writer) {
	lock.Lock()
	defer lock.Unlock()
	for _, tier := range tiers {
		tier.SetWriter(out)
	}
}

// SetTimeFormat sets the time format for the time stamp on a log line
// Uses the go standard library timeformat format.
func SetTimeFormat(format string) {
	lock.Lock()
	defer lock.Unlock()
	for _, tier := range tiers {
		tier.SetTimeFormat(NewTimeFormat(format))
	}
}

// SetFormat will set the logger to format all output.
// The format string
// MUST have a {TIME}, {TAG}, {MSG} string inside.
// For example: `{TIME} [{TAG}]:> {MSG}` will print logs of the form
// `10-21-1975 13:24:56 ERROR:> this is the message`
// Returns an error if an error occurs.
func SetFormat(format string) error {
	f := NewFormat(format)
	if f == nil {
		return fmt.Errorf(invaledTimeFormatFmt, format)
	}

	lock.Lock()
	defer lock.Unlock()
	for _, tier := range tiers {
		tier.SetFormat(f)
	}
	return nil
}

// Get will git the tier at idx
func Get(idx int) ITier {
	lock.RLock()
	defer lock.RUnlock()
	return tiers[idx]
}

// GetInfo gets the info tier
func GetInfo() ITier {
	return Get(infoIdx)
}

// Infof prints to output f formatted with values i.
// Uses the go standard library style format strings.
// Will add the assigned Info tag and color.
func Infof(f string, i ...interface{}) {
	GetInfo().Logf(f, i...)
}

// Info prints to output values i joined with a space.
// Will add the assigned Info tag and color.
func Info(i ...interface{}) {
	GetInfo().Log(i...)
}

// GetDebug gets the info tier
func GetDebug() ITier {
	return Get(debugIdx)
}

// Debugf prints to output f formatted with values i.
// Uses the go standard library style format strings.
// Will add the assigned Debug tag and color.
func Debugf(f string, i ...interface{}) {
	GetDebug().Logf(f, i...)
}

// Debug prints to output values i joined with a space.
// Will add the assigned Debug tag and color.
func Debug(i ...interface{}) {
	GetDebug().Log(i...)
}

// GetWarn gets the info tier
func GetWarn() ITier {
	return Get(warnIdx)
}

// Warnf prints to output f formatted with values i.
// Uses the go standard library style format strings.
// Will add the assigned Warn tag and color.
func Warnf(f string, i ...interface{}) {
	GetWarn().Logf(f, i...)
}

// Warn prints to output values i joined with a space.
// Will add the assigned Warn tag and color.
func Warn(i ...interface{}) {
	GetWarn().Log(i...)
}

// GetError gets the info tier
func GetError() ITier {
	return Get(errorIdx)
}

// Errorf prints to output f formatted with values i.
// Uses the go standard library style format strings.
// Will add the assigned Error tag and color.
func Errorf(f string, i ...interface{}) {
	GetError().Logf(f, i...)
}

// Error prints to output values i joined with a space.
// Will add the assigned Error tag and color.
func Error(i ...interface{}) {
	GetError().Log(i...)
}

// GetFatal gets the info tier
func GetFatal() ITier {
	return Get(fatalIdx)
}

// Fatalf prints to output f formatted with values i.
// Uses the go standard library style format strings.
// Will add the assigned Fatal tag and color.
func Fatalf(exit func(int), f string, i ...interface{}) {
	GetFatal().Logf(f, i...)
	exit(1)
}

// Fatal prints to output values i joined with a space.
// Will add the assigned Fatal tag and color.
func Fatal(exit func(int), i ...interface{}) {
	GetFatal().Log(i...)
	exit(1)
}
