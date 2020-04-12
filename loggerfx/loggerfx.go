package loggerfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ProvideLogger to fx
func ProvideLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()

	return slogger
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(ProvideLogger),
)
