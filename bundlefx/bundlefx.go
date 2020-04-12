package bundlefx

import (
	"context"
	"net/http"

	"github.com/preslavmihaylov/fxappexample/configfx"
	"github.com/preslavmihaylov/fxappexample/httpfx"
	"github.com/preslavmihaylov/fxappexample/loggerfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func registerHooks(
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger, cfg *configfx.Config, mux *http.ServeMux,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info("Listening on ", cfg.ApplicationConfig.Address)
				go http.ListenAndServe(cfg.ApplicationConfig.Address, mux)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}

// Module provided to fx
var Module = fx.Options(
	configfx.Module,
	loggerfx.Module,
	httpfx.Module,
	fx.Invoke(registerHooks),
)
