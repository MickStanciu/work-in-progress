package util

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if logger != nil {
		return logger
	}

	l, _ := zap.NewProduction()
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {

		}
	}(l)
	logger = l.Sugar()
	return logger
}
