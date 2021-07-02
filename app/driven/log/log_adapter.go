package logbook

import (
	"io/ioutil"
	"log"

	"github.com/google/logger"
)

//NewLogger creates a log struct
//log only on stdout
func NewLogger() *logger.Logger {
	logbook := logger.Init("logger", true, false, ioutil.Discard)
	logger.SetFlags(log.Ldate)

	defer logbook.Close()

	return logbook
}
