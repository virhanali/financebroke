package utils

import (
	"fmt"
	"log"
	"runtime"
)

type Logger struct{}

func GetLogger() *Logger {
	return &Logger{}
}

type ErrorWithContext struct {
	Err      error
	Context  string
	Function string
	File     string
	Line     int
}

func (e *ErrorWithContext) Error() string {
	return fmt.Sprintf("[%s] %s: %v (%s:%d)", e.Context, e.Function, e.Err, e.File, e.Line)
}

func LogError(ctx string, err error) error {
	if err == nil {
		return nil
	}

	// Get caller info
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		log.Printf("[ERROR] [%s] %v (%s:%d)", ctx, err, file, line)
		return err
	}

	funcName := fn.Name()

	// Log detailed error
	log.Printf("[ERROR] [%s] %s: %v (%s:%d)", ctx, funcName, err, file, line)

	return &ErrorWithContext{
		Err:      err,
		Context:  ctx,
		Function: funcName,
		File:     file,
		Line:     line,
	}
}

func (l *Logger) Info(message string, fields ...map[string]interface{}) {
	// Get caller info
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		log.Printf("[INFO] %s", message)
		return
	}

	funcName := fn.Name()

	// Log with fields if any
	if len(fields) > 0 && len(fields[0]) > 0 {
		log.Printf("[INFO] %s: %s %+v", funcName, message, fields[0])
	} else {
		log.Printf("[INFO] %s: %s (%s:%d)", funcName, message, file, line)
	}
}

func WrapError(ctx string, err error, message string) error {
	if err == nil {
		return nil
	}

	wrapped := fmt.Errorf("%s: %w", message, err)
	return LogError(ctx, wrapped)
}