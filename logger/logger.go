package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	options struct {
		level string
	}

	OptionFunc func(*options)
)

// ログのレベルを設定する
func WithLevel(level string) OptionFunc {
	return func(opts *options) {
		opts.level = level
	}
}

func NewLogger(opts ...OptionFunc) (*zap.Logger, error) {
	dops := &options{
		level: "info",
	}
	for _, op := range opts {
		op(dops)
	}
	level := new(zapcore.Level)
	if err := level.Set(dops.level); err != nil {
		fmt.Println(dops.level)
		return nil, err
	}
	return NewProductionConfig(*level).Build()
}

// Development設定
func NewDevelopmentConfig(level zapcore.Level) zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// Productoin設定
func NewProductionConfig(level zapcore.Level) zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}
