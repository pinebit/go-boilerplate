package logger

import "go.uber.org/zap"

func NewNoopLogger() Logger {
	return &logger{
		zap: zap.NewNop().Sugar(),
	}
}
