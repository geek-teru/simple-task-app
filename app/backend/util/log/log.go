package log

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type option struct {
	initFields map[string]interface{}
}

type Option func(*option)

func WithCommandName(commandName string) Option {
	return func(o *option) {
		o.initFields = map[string]interface{}{"command": commandName}
	}
}

func New(logLevel string, opts ...Option) (*zap.Logger, error) {
	var l zapcore.Level
	if err := l.Set(logLevel); err != nil {
		fmt.Fprintf(os.Stderr, "the value of LOG_LEVEL is invalid: err=%v. use default log level: %q\n", err, "info")
		l = zapcore.InfoLevel
	}

	o := &option{
		initFields: map[string]interface{}{},
	}
	for _, opt := range opts {
		opt(o)
	}

	conf := newConf(l)
	conf.InitialFields = o.initFields
	logger, err := conf.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func newConf(logLevel zapcore.Level) zap.Config {
	c := zap.Config{
		Level:             zap.NewAtomicLevelAt(logLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "timestamp",
			NameKey:        "name",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	return c
}

func Sync(l *zap.Logger) {
	if l != nil {
		if err := l.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) && !errors.Is(err, syscall.EINVAL) && !errors.Is(err, os.ErrInvalid) {
			fmt.Fprintf(os.Stderr, "failed to zap.Logger.Sync: %v", err)
		}
	}
}
