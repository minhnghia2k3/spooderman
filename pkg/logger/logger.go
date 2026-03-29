package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/rs/zerolog"
)

type LogLevel = zerolog.Level

const (
	DEBUG = zerolog.DebugLevel
	INFO  = zerolog.InfoLevel
	WARN  = zerolog.WarnLevel
	ERROR = zerolog.ErrorLevel
	FATAL = zerolog.FatalLevel
)

var logLevelName = map[LogLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

var (
	currentLevel = DEBUG
)

var (
	once       sync.Once
	logger     zerolog.Logger
	fileLogger zerolog.Logger
)

const (
	locUnknown = "<unknown>"
	Component  = "component"
)

func init() {
	once.Do(func() {
		zerolog.SetGlobalLevel(currentLevel)

		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "15:04:03",
			// FormatFieldValue: formatFieldValue, // TODO: implement format field value if needed
		}

		logger = zerolog.New(consoleWriter).With().Timestamp().Caller().Logger()
		fileLogger = zerolog.Logger{}
	})
}

func parseLevel(s string) (LogLevel, bool) {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "debug":
		return DEBUG, true
	case "info":
		return INFO, true
	case "warn":
		return WARN, true
	case "error":
		return ERROR, true
	case "fatal":
		return FATAL, true
	default:
		return INFO, false
	}
}

// SetLevelFromString sets log level from a string value.
// If string is empty or not regconized log level -> the current is kept.
func SetLevelFromString(s string) {
	if s == "" {
		return
	}

	if level, ok := parseLevel(s); ok {
		zerolog.SetGlobalLevel(level)
	}
}

func getCallerSkip() (int, string) {
	for i := 2; i < 15; i++ {
		pc, file, _, ok := runtime.Caller(i)
		if !ok {
			continue
		}

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		// bypass runtime.*
		if strings.Contains(fn.Name(), "runtime.") {
			continue
		}

		return i - 1, getPackageFromFile(file)
	}
	return 3, locUnknown
}

// getPackageFromFile parse the file path to get caller pkg
// e.g.  "/home/usr/spooderman/pkg/config/config.go" -> config
func getPackageFromFile(file string) string {
	dir := filepath.Dir(file)           // /home/usr/spooderman/pkg/config
	importPath := filepath.ToSlash(dir) // converts \ to / on Windows

	parts := strings.Split(importPath, "/")
	if len(parts) == 0 {
		return locUnknown
	}

	pkg := parts[len(parts)-1]
	if pkg == "." {
		return "<main>"
	}

	return pkg
}

func getEvent(logger zerolog.Logger, level LogLevel) *zerolog.Event {
	switch level {
	case zerolog.DebugLevel:
		return logger.Debug()
	case zerolog.InfoLevel:
		return logger.Info()
	case zerolog.WarnLevel:
		return logger.Warn()
	case zerolog.ErrorLevel:
		return logger.Error()
	case zerolog.FatalLevel:
		return logger.Fatal()
	default:
		return logger.Info()
	}
}

func logMessage(level LogLevel, component string, message string, fields map[string]any) {
	if level < currentLevel {
		return
	}

	skip, pkg := getCallerSkip()

	event := getEvent(logger, level)

	if component == "" {
		component = pkg
	}

	event.Str(Component, component)
	// appendFields(event, fields)

	event.CallerSkipFrame(skip).Msg(message)

	// TODO: log to the file if needed

	if level == FATAL {
		os.Exit(1)
	}
}

func Debug(message string) {
	logMessage(DEBUG, "", message, nil)
}

func Debugf(message string, ss ...any) {
	logMessage(DEBUG, "", fmt.Sprintf(message, ss...), nil)
}

func Warn(message string) {
	logMessage(WARN, "", message, nil)
}

func Warnf(message string, ss ...any) {
	logMessage(WARN, "", fmt.Sprintf(message, ss...), nil)
}

func WarnF(message string, fields map[string]any) {
	logMessage(WARN, "", message, fields)
}

func Error(message string) {
	logMessage(ERROR, "", message, nil)
}

func Errorf(message string, ss ...any) {
	logMessage(ERROR, "", fmt.Sprintf(message, ss...), nil)
}

func Fatal(message string) {
	logMessage(FATAL, "", message, nil)
}

func Fatalf(message string, ss ...any) {
	logMessage(FATAL, "", fmt.Sprintf(message, ss...), nil)
}
