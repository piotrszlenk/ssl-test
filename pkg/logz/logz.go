package logz

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// LogHandler stores pointers to various log level loggers
type LogHandler struct {
	Debug   *log.Logger
	Error   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
}

type LoggerConfig struct {
	Debug bool
}

// SetHandles sets log handles for each log level
func (e *LogHandler) SetHandles(traceHandle io.Writer, errorHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer) {
	e.Debug = log.New(traceHandle, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	e.Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	e.Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	e.Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// GetInstance returns LogHandler instance.
func GetInstance(lc LoggerConfig) *LogHandler {
	var logger *LogHandler = new(LogHandler)
	if lc.Debug {
		logger.SetHandles(os.Stdout, os.Stderr, os.Stdout, os.Stdout)
	} else {
		logger.SetHandles(ioutil.Discard, os.Stderr, os.Stdout, os.Stdout)
	}
	return logger
}
