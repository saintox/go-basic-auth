package logger

import (
	"context"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
)

type Config struct {
	EnableDebugMode      bool
	EnableConsoleLogging bool
	EncodeLogsAsJson     bool
	EnableFileLogging    bool
	FileDir              string
	FileName             string
	MaxLogSize           int
	MaxBackups           int
	MaxAge               int
}

type Logger struct {
	logger *zerolog.Logger
}

func Configure(config Config) *Logger {
	var writers []io.Writer
	logLevel := zerolog.InfoLevel

	if config.EnableConsoleLogging {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if config.EnableFileLogging {
		writers = append(writers, newRollingFile(config))
	}

	if config.EnableDebugMode {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	mw := io.MultiWriter(writers...)

	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", config.EnableFileLogging).
		Bool("jsonLogOutput", config.EncodeLogsAsJson).
		Str("logDirectory", config.FileDir).
		Str("fileName", config.FileName).
		Int("maxSizeMB", config.MaxLogSize).
		Int("maxAgeInDays", config.MaxAge).
		Msg("configuration loaded")

	return &Logger{
		logger: &logger,
	}
}

func newRollingFile(config Config) io.Writer {
	return &lumberjack.Logger{
		Filename:   path.Join(config.FileDir, config.FileName),
		MaxSize:    config.MaxLogSize, // n megabytes
		MaxAge:     config.MaxAge,     // n days
		MaxBackups: config.MaxBackups, // n files
		LocalTime:  false,
		Compress:   false,
	}
}

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
