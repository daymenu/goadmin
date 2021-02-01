package pkg

import (
	"github.com/daymenu/goadmin/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
}

// New new
func New() {
	conf, err := config.New().Load()
	if err != nil {
		panic(err)
	}
	c := conf.GetLog()
	jackLog := lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    c.MaxSize,
		MaxAge:     c.MaxAge,
		MaxBackups: c.MaxBackups,
	}

	encoderConfig := zapcore.EncoderConfig{}
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&jackLog)),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)

	development := zap.Development()

	logger := zap.New(core, zap.AddCaller(), development)

	zap.ReplaceGlobals(logger)
}
