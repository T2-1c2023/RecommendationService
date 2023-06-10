package utilities

import (
	"github.com/fatih/color"
)

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
)

type ILogger interface {
	LogDebug(message string)
	LogInfo(message string)
	LogWarn(message string)
	LogError(err error)
}

type Logger struct {
	Level LogLevel
}

func NewLogger(logLevel string) Logger {
	logger := Logger{Level: LogLevel(logLevel)}
	switch logger.Level {
	case Debug, Info, Warn, Error:
		return logger
	default:
		logger.Level = Debug
		return logger
	}
}

func (logger *Logger) LogDebug(message string) {
	if logger.Level == Debug {
		color.Cyan("[DEBUG]\t\t" + message)
	}
}

func (logger *Logger) LogInfo(message string) {
	if logger.Level == Info || logger.Level == Debug {
		color.White("[INFO]\t\t" + message)
	}
}

func (logger *Logger) LogWarn(message string) {
	if logger.Level != Error {
		color.Yellow("[WARN]\t\t" + message)
	}
}

func (logger *Logger) LogError(err error) {
	if logger.Level != Error {
		color.Red("[ERROR]\t\t" + err.Error())
	}
}
