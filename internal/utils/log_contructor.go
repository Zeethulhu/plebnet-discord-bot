package utils

import (
	"log"
	"os"
)

func NewLogger(pkg string) *log.Logger {
	return log.New(os.Stdout, "["+pkg+"] ", log.LstdFlags)
}
