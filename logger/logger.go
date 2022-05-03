package logger

import (
	"aspire/config"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

var ZeroLogger *Logger

var SubLogger zerolog.Logger

var fileName string

var fileMap = make(map[string]bool)

func InitLogger() {
	ZeroLogger = New("")
}

func New(fileName string) *Logger {
	var writers []io.Writer
	logConfig := config.Conf.Logger

	if logConfig.Isconsoleenabled {
		writers = append(writers, os.Stdout)
	}
	if logConfig.Isfileenabled {
		writers = append(writers, newRollingFile(logConfig))
	}
	if logConfig.Isstderrenabled {
		writers = append(writers, os.Stderr)
	}
	mw := io.MultiWriter(writers...)

	logLevel := zerolog.InfoLevel
	if logConfig.Level == "debug" {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	zerolog.ErrorStackMarshaler = getErrorStackMarshaller()
	logger := zerolog.New(mw).With().Timestamp().Logger()
	return &Logger{logger: &logger}
}

func getErrorStackMarshaller() func(err error) interface{} {
	return func(err error) interface{} {
		return string(debug.Stack())
	}
}

func UpdateLogger() *Logger {
	var writers []io.Writer

	logConfig := config.Conf.Logger

	if logConfig.Isconsoleenabled {
		writers = append(writers, os.Stderr)
	}
	if logConfig.Isfileenabled {
		writers = append(writers, newRollingFile(logConfig))
	}
	if logConfig.Isstderrenabled {
		writers = append(writers, os.Stderr)
	}
	mw := io.MultiWriter(writers...)

	logLevel := zerolog.InfoLevel
	if logConfig.Level == "debug" {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(mw).With().Timestamp().Logger()
	ZeroLogger.logger = &logger
	return ZeroLogger
}

func ZeroLog() *Logger {

	t := time.Now()
	fileName = "discoveryservice_" + strconv.Itoa(t.Day()) + "_" + t.Month().String() + "_" + strconv.Itoa(t.Year()) + "_" + strconv.Itoa(t.Hour()) + "_" + strconv.Itoa(t.Minute()) + ".log"

	if _, ok := fileMap[fileName]; ok {
		return ZeroLogger
	}
	return UpdateLogger()
}

func newRollingFile(logConfig config.LoggerConfig) io.Writer {
	t := time.Now()
	fileName = "discoveryservice_" + strconv.Itoa(t.Day()) + "_" + t.Month().String() + "_" + strconv.Itoa(t.Year()) + "_" + strconv.Itoa(t.Hour()) + "_" + strconv.Itoa(t.Minute()) + ".log"

	fileMap[fileName] = true

	tempFile, err := os.Create(filepath.Join(logConfig.Logfilepath, fileName))

	if err != nil {
		log.Print(err)

	}

	return tempFile

}

func NewConsole(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &Logger{logger: &logger}
}

// Output duplicates the global logger and sets w as its output.
func (l *Logger) Output(w io.Writer) zerolog.Logger {
	return l.logger.Output(w)
}

// With creates a child logger with the field added to its context.
func (l *Logger) With() zerolog.Context {
	return l.logger.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func (l *Logger) Level(level zerolog.Level) zerolog.Logger {
	return l.logger.Level(level)
}

// Sample returns a logger with the s sampler.
func (l *Logger) Sample(s zerolog.Sampler) zerolog.Logger {
	return l.logger.Sample(s)
}

// Hook returns a logger with the h Hook.
func (l *Logger) Hook(h zerolog.Hook) zerolog.Logger {
	return l.logger.Hook(h)
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Panic() *zerolog.Event {
	return l.logger.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) WithLevel(level zerolog.Level) *zerolog.Event {
	return l.logger.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Log() *zerolog.Event {
	return l.logger.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) {
	l.logger.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func (l *Logger) Ctx(ctx context.Context) *Logger {
	return &Logger{logger: zerolog.Ctx(ctx)}
}
