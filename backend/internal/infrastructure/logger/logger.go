package logger

import (
	"context"
	"fmt"
	go_log "log"
	"log/slog"
	"os"
	"strings"
	"sync"
)

var (
	once sync.Once
)

func Init(logLevel string) {
	once.Do(func() {
		// note: config default one in case other parts of app don't use logger but slog directly
		go_log.SetFlags(go_log.LstdFlags | go_log.Lmicroseconds | go_log.LUTC)
		go_log.SetOutput(os.Stdout)

		var level slog.Level
		switch strings.ToLower(logLevel) {
		case "debug":
			level = slog.LevelDebug
		case "info":
			level = slog.LevelInfo
		case "warn":
			level = slog.LevelWarn
		case "error":
			level = slog.LevelError
		default:
			level = slog.LevelInfo
		}

		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})))
	})
}

func Discard() {
	once.Do(func() {
		slog.SetDefault(slog.New(slog.DiscardHandler))
	})
}

func Default() *slog.Logger {
	return slog.Default()
}

// ErrAttr creates a slog.Attr for an error, ensuring consistent error logging.
func ErrAttr(err error) slog.Attr {
	return slog.String("error", err.Error())
}

func log(level slog.Level, msg string, args ...any) {
	slog.Log(context.Background(), level, msg, args...)
}

func Info(msg string, args ...any) {
	log(slog.LevelInfo, msg, args...)
}

func Error(msg string, args ...any) {
	if len(args) == 1 {
		if err, ok := args[0].(error); ok {
			log(slog.LevelError, msg, ErrAttr(err))
			return
		}
	}
	log(slog.LevelError, msg, args...)
}

func CombineErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	errorStrings := make([]string, len(errs))
	for i, err := range errs {
		errorStrings[i] = err.Error()
	}
	return fmt.Errorf("multiple errors occurred: %s", strings.Join(errorStrings, "; "))
}

func Debug(msg string, args ...any) {
	log(slog.LevelDebug, msg, args...)
}

func Warn(msg string, args ...any) {
	log(slog.LevelWarn, msg, args...)
}
