package commands

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[Discord] ", log.LstdFlags)