package logz

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var logger *LogHandler
var initLog sync.Once

// LogHandler stores pointers to various log level loggers
type LogHandler struct {
	Debug     *log.Logger
	Error     *log.Logger
	Info      *log.Logger
	Warning   *log.Logger
	DebugMode bool
}

// SetHandles sets log handles for each log level
func (e *LogHandler) SetHandles(traceHandle io.Writer, errorHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer) {
	e.Debug = log.New(traceHandle, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	e.Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	e.Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	e.Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// GetInstance returns LogHandler instance.
func InitLog() *LogHandler {
	initLog.Do(func() {
		logger = new(LogHandler)
		logger.SetHandles(ioutil.Discard, os.Stderr, os.Stdout, os.Stdout)
	})
	return logger
}

func InitDebugLog() *LogHandler {
	initLog.Do(func() {
		logger = new(LogHandler)
		logger.SetHandles(os.Stdout, os.Stderr, os.Stdout, os.Stdout)
	})
	return logger
}

func Logger() *LogHandler {
	return logger
}
