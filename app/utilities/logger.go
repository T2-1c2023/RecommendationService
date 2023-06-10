package utilities

import (
	"os"

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
	Level      LogLevel
	LogFile    *os.File
	DebugColor *color.Color
	InfoColor  *color.Color
	WarnColor  *color.Color
	ErrorColor *color.Color
}

func NewLogger(logLevel string, logFile *os.File) Logger {
	logger := Logger{
		Level:      LogLevel(logLevel),
		LogFile:    logFile,
		DebugColor: color.New(color.BgCyan),
		InfoColor:  color.New(color.BgWhite),
		WarnColor:  color.New(color.BgYellow),
		ErrorColor: color.New(color.BgRed),
	}
	switch logger.Level {
	case Debug, Info, Warn, Error:
		return logger
	default:
		logger.Level = Debug
		return logger
	}
}

func (logger *Logger) CloseFile() {
	logger.LogFile.Close()
}

func (logger *Logger) log(message string, logColor *color.Color) {
	logColor.Fprintln(logger.LogFile, message)
	logColor.Println(message)
}

func (logger *Logger) LogDebug(message string) {
	if logger.Level == Debug {
		logger.log("[DEBUG]\t\t"+message, logger.DebugColor)
	}
}

func (logger *Logger) LogInfo(message string) {
	if logger.Level == Info || logger.Level == Debug {
		logger.log("[INFO]\t\t"+message, logger.InfoColor)
	}
}

func (logger *Logger) LogWarn(message string) {
	if logger.Level != Error {
		logger.log("[WARN]\t\t"+message, logger.WarnColor)
	}
}

func (logger *Logger) LogError(err error) {
	if logger.Level != Error {
		logger.log("[ERROR]\t\t"+err.Error(), logger.ErrorColor)
	}
}
