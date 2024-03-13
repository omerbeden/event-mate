package pkg

import (
	"sync"

	"go.uber.org/zap"
)

type contextKey int

const (
	LoggerKey contextKey = iota
)

var (
	logger *zap.SugaredLogger
	once   sync.Once
)

func Logger() *zap.SugaredLogger {
	once.Do(func() {
		l, _ := zap.NewProduction(zap.AddCaller(), zap.AddStacktrace(zap.InfoLevel))
		logger = l.Sugar()
	})
	return logger
}
