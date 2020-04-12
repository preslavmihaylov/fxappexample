package main

import (
	"github.com/preslavmihaylov/fxappexample/bundlefx"
	"github.com/preslavmihaylov/fxappexample/httphandler"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bundlefx.Module,
		fx.Invoke(httphandler.New),
	).Run()
}
