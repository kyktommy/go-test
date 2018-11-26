package logger

import (
	"flag"

	"github.com/google/logger"
)

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func Log2(s string) {
	logger.Infof("I'm google logger, " + s)
}
