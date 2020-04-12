package httpfx

import (
	"net/http"

	"go.uber.org/fx"
)

// Module provided to fx
var Module = fx.Options(
	fx.Provide(http.NewServeMux),
)
