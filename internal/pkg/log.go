package pkg

import (
	"os"

	"github.com/daymenu/goadmin/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger logger
type Logger struct {
	opts logOptions
	*zap.Logger
}

type logOptions struct {
	conf  config.Log
	level zapcore.Level
}

var defaultLogOptions = logOptions{
	conf: config.Log{
		Filename:   "./app.log",
		MaxSize:    1024,
		MaxAge:     3,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   false,
	},
}

// LogOption A log option
type LogOption interface {
	apply(*logOptions)
}

type funcLogOption struct {
	f func(*logOptions)
}

func (fdo *funcLogOption) apply(do *logOptions) {
	fdo.f(do)
}

func newFuncLogOption(f func(*logOptions)) *funcLogOption {
	return &funcLogOption{
		f: f,
	}
}

// Conf set log config
func Conf(c config.Log) LogOption {
	return newFuncLogOption(func(l *logOptions) {
		l.conf = c
	})
}

// NewLog new
func NewLog(opt ...LogOption) *Logger {
	// 设置选项
	opts := defaultLogOptions
	for _, o := range opt {
		o.apply(&opts)
	}

	// 建立日志归档
	jackLog := &lumberjack.Logger{
		Filename:   opts.conf.Filename,
		MaxSize:    opts.conf.MaxSize,
		MaxAge:     opts.conf.MaxAge,
		MaxBackups: opts.conf.MaxBackups,
	}

	encoderConfig := defaultEncoderConfig()
	encoder := zapcore.NewJSONEncoder(*encoderConfig)
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(jackLog)),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)

	development := zap.Development()

	logger := &Logger{
		opts: opts,
	}
	logger.Logger = zap.New(core, zap.AddCaller(), development)
	return logger
}

func defaultEncoderConfig() *zapcore.EncoderConfig {
	return &zapcore.EncoderConfig{
		LevelKey:       "level",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		TimeKey:        "time",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		NameKey:        "name",
		EncodeName:     zapcore.FullNameEncoder,
		CallerKey:      "caller",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		MessageKey:     "msg",
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
}
