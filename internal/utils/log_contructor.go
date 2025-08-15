package utils

import (
	"io"
	"log"
	"os"
)

var verbose bool

// SetVerbose enables or disables verbose logging for all loggers created via
// this package.
func SetVerbose(v bool) {
	verbose = v
}

type verboseWriter struct{}

func (verboseWriter) Write(p []byte) (int, error) {
	if verbose {
		return os.Stdout.Write(p)
	}
	return io.Discard.Write(p)
}

// NewLogger returns a logger that writes output only when verbose mode is
// enabled via SetVerbose.
func NewLogger(pkg string) *log.Logger {
	return log.New(verboseWriter{}, "["+pkg+"] ", log.LstdFlags)
}
