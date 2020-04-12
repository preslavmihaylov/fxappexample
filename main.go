package main

import (
	"io/ioutil"
	"net/http"

	"github.com/preslavmihaylov/fxappexample/httphandler"
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

func main() {
	conf := &Config{}
	data, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	slogger := logger.Sugar()

	mux := http.NewServeMux()
	httphandler.New(mux, slogger)

	http.ListenAndServe(conf.ApplicationConfig.Address, mux)
}
