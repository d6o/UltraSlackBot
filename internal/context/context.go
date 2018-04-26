package context

import (
	"context"
	"errors"
	"log"
)

type indexContext int

const (
	outLoggerKey indexContext = iota
	errLoggerKey
	verbosityKey
)

func WithOutLogger(ctx context.Context, logger *log.Logger) context.Context {
	return context.WithValue(ctx, outLoggerKey, logger)
}

func WithErrLogger(ctx context.Context, logger *log.Logger) context.Context {
	return context.WithValue(ctx, errLoggerKey, logger)
}

func WithVerbosity(ctx context.Context, verbosity bool) context.Context {
	return context.WithValue(ctx, errLoggerKey, verbosity)
}

func OutLogger(ctx context.Context) *log.Logger {
	logger, _ := outLogger(ctx)
	return logger
}

func ErrLogger(ctx context.Context) *log.Logger {
	logger, _ := errLogger(ctx)
	return logger
}

func Verbosity(ctx context.Context) bool {
	val, _ := verbosity(ctx)
	return val
}

func outLogger(ctx context.Context) (*log.Logger, error) {
	outLogger, ok := ctx.Value(outLoggerKey).(*log.Logger)
	if !ok {
		return &log.Logger{}, errors.New("no OutLogger in the context")
	}

	return outLogger, nil
}

func errLogger(ctx context.Context) (*log.Logger, error) {
	errLogger, ok := ctx.Value(errLoggerKey).(*log.Logger)
	if !ok {
		return &log.Logger{}, errors.New("no ErrLogger in the context")
	}

	return errLogger, nil
}

func verbosity(ctx context.Context) (bool, error) {
	val, ok := ctx.Value(verbosityKey).(bool)
	if !ok {
		return false, errors.New("no ErrLogger in the context")
	}

	return val, nil
}
