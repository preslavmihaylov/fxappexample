package configfx

import (
	"io/ioutil"

	"go.uber.org/fx"
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

// Module provided to fx
var Module = fx.Options(
	fx.Provide(ProvideConfig),
)
