package main

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/preslavmihaylov/fxappexample/httphandler"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// ApplicationConfig ...
type ApplicationConfig struct {
	Address string `yaml:"address"`
}

// Config ...
type Config struct {
	ApplicationConfig `yaml:"application"`
}

// ProvideConfig provides the standard configuration to fx
func ProvideConfig() *Config {
	conf := Config{}
	data, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}

// ProvideLogger to fx
func ProvideLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()

	return slogger
}

func main() {
	fx.New(
		fx.Provide(ProvideConfig),
		fx.Provide(ProvideLogger),
		fx.Provide(http.NewServeMux),
		fx.Provide(httphandler.New),
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger, cfg *Config, handler *httphandler.Handler,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go handler.ListenAndServe(cfg.ApplicationConfig.Address)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}
