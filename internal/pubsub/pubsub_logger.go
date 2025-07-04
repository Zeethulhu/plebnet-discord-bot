package pubsub

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[PubSub] ", log.LstdFlags)
