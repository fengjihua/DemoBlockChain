package lib

import (
	"fmt"
	"io"
	"os"

	"github.com/op/go-logging"
)

/*
Log :Global Log Instance
*/
var Log = logging.MustGetLogger("example")

func init() {
	// fmt.Println("Log init")

	// Example format string. Everything except the message has a custom color
	// which is dependent on the log level. Many fields have a custom output
	// formatting too, eg. the time returns the hour down to the milli second.
	var formatStd = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	var formatShort = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	f, err := os.Create("logs/err.log")
	if err != nil {
		fmt.Println("Log init error:", err)
	}
	// multiWriter := io.MultiWriter(f, os.Stdout)
	multiWriter := io.MultiWriter(f)

	backend1 := logging.NewLogBackend(multiWriter, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend1Formatter := logging.NewBackendFormatter(backend1, formatShort)
	backend2Formatter := logging.NewBackendFormatter(backend2, formatStd)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1Formatter)
	backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)

	Log.Debugf("debug %s", "hello")
	Log.Info("info")
	Log.Notice("notice")
	Log.Warning("warning")
	Log.Error("err")
	Log.Critical("crit")
}
