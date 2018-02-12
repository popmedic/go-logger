package tty

import "fmt"

// TTY is an int holding number escape value for a color on the terminal.
type TTY int

const (
	defaultInfoTermColor  TTY = Reset
	defaultDebugTermColor TTY = FgBlue
	defaultWarnTermColor  TTY = FgYellow
	defaultErrorTermColor TTY = FgRed
	defaultFatalTermColor TTY = FgGreen
)

// Font Types
const (
	Reset TTY = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground Colors
const (
	FgBlack TTY = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground High Colors
const (
	FgHiBlack TTY = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background Colors
const (
	BgBlack TTY = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background High Colors
const (
	BgHiBlack TTY = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

const (
	unset  = "\x1b[0m"
	format = "\x1b[%d;1m"
)

// String returns the escape for a color as a string.
func (c TTY) String() string {
	if c == Reset {
		return unset
	}
	return fmt.Sprintf(format, c)
}
