package examples

import (
	"os"

	"github.com/popmedic/go-color/colorize/tty"
	"github.com/popmedic/go-logger/log"
)

// Run1 runs the first example
func Run1() {
	// Show the defaults...
	log.Info("hello,")
	log.Infof("this %q", "is")
	log.Info("a", "test")
	log.Debug("hello,")
	log.Debugf("this %q", "is")
	log.Debug("a", "test")
	log.Warn("hello,")
	log.Warnf("this %q", "is")
	log.Warn("a", "test")
	log.Error("hello,")
	log.Errorf("this %q", "is")
	log.Error("a", "test")
	log.Fatal(func(int) {}, "good bye")
	// Change it up...
	log.GetInfo().SetColor(log.NewColor(tty.FgHiBlue(), tty.Underline()))
	log.GetDebug().SetColor(log.NewColor(tty.FgMagenta()))
	log.GetWarn().SetColor(log.NewColor(tty.BgYellow().Add(tty.FgHiBlue(), tty.Underline())))
	log.GetError().SetColor(log.NewColor(tty.FgCyan()))
	log.GetFatal().SetColor(log.NewColor(tty.FgHiGreen()))
	log.GetInfo().SetTag(log.NewTag("information"))
	log.GetDebug().SetTag(log.NewTag("   debug   "))
	log.GetWarn().SetTag(log.NewTag("  warning  "))
	log.GetError().SetTag(log.NewTag(" erroring! "))
	log.GetFatal().SetTag(log.NewTag("!fatality!!"))

	log.SetFormat("<{TAG}> : {MSG} : [{TIME}]")
	log.SetTimeFormat("Mon Jan _2 15:04:05 2006")

	log.Info("hello,")
	log.Infof("this %q", "is")
	log.Info("a", "test")
	log.Debug("hello,")
	log.Debugf("this %q", "is")
	log.Debug("a", "test")
	log.Warn("hello,")
	log.Warnf("this %q", "is")
	log.Warn("a", "test")
	log.Error("hello,")
	log.Errorf("this %q", "is")
	log.Error("a", "test")
	log.Fatal(os.Exit, "good bye")
}
