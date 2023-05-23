package logger

import "go.uber.org/zap"

type Logger interface {
	Named(name string) Logger

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})

	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})

	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	Zap() *zap.Logger
}

type logger struct {
	zap *zap.SugaredLogger
}

var _ Logger = (Logger)(&logger{})

func NewLogger(devMode bool) (Logger, error) {
	var zl *zap.Logger
	var err error
	if devMode {
		zl, err = zap.NewDevelopment()
	} else {
		zl, err = zap.NewProduction()
	}
	if err != nil {
		return nil, err
	}
	return &logger{zap: zl.Sugar()}, nil
}

func (l logger) Named(name string) Logger {
	return &logger{zap: l.zap.Named(name)}
}

func (l logger) Debug(args ...interface{}) {
	l.zap.Debug(args...)
}

func (l logger) Info(args ...interface{}) {
	l.zap.Info(args...)
}

func (l logger) Warn(args ...interface{}) {
	l.zap.Warn(args...)
}

func (l logger) Error(args ...interface{}) {
	l.zap.Error(args...)
}

func (l logger) Panic(args ...interface{}) {
	l.zap.Panic(args...)
}

func (l logger) Fatal(args ...interface{}) {
	l.zap.Fatal(args...)
}

func (l logger) Debugf(template string, args ...interface{}) {
	l.zap.Debugf(template, args...)
}

func (l logger) Infof(template string, args ...interface{}) {
	l.zap.Infof(template, args...)
}

func (l logger) Warnf(template string, args ...interface{}) {
	l.zap.Warnf(template, args...)
}

func (l logger) Errorf(template string, args ...interface{}) {
	l.zap.Errorf(template, args...)
}

func (l logger) Panicf(template string, args ...interface{}) {
	l.zap.Panicf(template, args...)
}

func (l logger) Fatalf(template string, args ...interface{}) {
	l.zap.Fatalf(template, args...)
}

func (l logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.zap.Debugw(msg, keysAndValues...)
}

func (l logger) Infow(msg string, keysAndValues ...interface{}) {
	l.zap.Infow(msg, keysAndValues...)
}

func (l logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.zap.Warnw(msg, keysAndValues...)
}

func (l logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.zap.Errorw(msg, keysAndValues...)
}

func (l logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.zap.Panicw(msg, keysAndValues...)
}

func (l logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.zap.Fatalw(msg, keysAndValues...)
}

func (l logger) Zap() *zap.Logger {
	return l.zap.Desugar()
}
