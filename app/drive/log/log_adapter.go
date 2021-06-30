package logbook

import (
	"io/ioutil"
	"log"

	"github.com/google/logger"
)

func NewLogger() *logger.Logger {
	logbook := logger.Init("logger", true, false, ioutil.Discard)
	logger.SetFlags(log.Ldate)

	defer logbook.Close()

	return logbook
}
