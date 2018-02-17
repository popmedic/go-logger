package examples

import (
	"math/rand"
	"time"

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
	log.Fatal(func(int) {}, "good bye")
}

func Run2() {
	log.GetInfo().SetTag(log.NewTag("INFO"))
	log.GetDebug().SetTag(log.NewTag("DBUG"))
	log.GetWarn().SetTag(log.NewTag("WARN"))
	log.GetError().SetTag(log.NewTag("EROR"))
	log.GetFatal().SetTag(log.NewTag("FATL"))

	log.SetFormat("[{TAG}] {TIME} -> {MSG}")
	log.SetTimeFormat("01-02-2006 15:04:05")

	log.SetHTMLStatus(true, ":8180")
	for {
		rand.Seed(time.Now().Unix())
		idx := rand.Int()
		switch idx % 5 {
		case 0:
			log.Info("A info number: ", idx)
		case 1:
			log.Debug("A debug number: ", idx)
		case 2:
			log.Warn("A warn number: ", idx)
		case 3:
			log.Error("A error number: ", idx)
		case 4:
			log.Fatal(func(int) {}, "A fatal number: ", idx)
		}
		time.Sleep(time.Second)
	}
}
