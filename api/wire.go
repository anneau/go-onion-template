//go:build wireinject

package api

import (
	"github.com/anneau/go-template/api/controller"
	"github.com/anneau/go-template/api/router"
	"github.com/google/wire"
)

func InitializeRouter() (*router.Router, error) {
	panic(
		wire.Build(
			wire.Struct(new(router.Router), "*"),
			InitializeHealthController,
		),
	)
}

func InitializeHealthController() (*controller.HealthController, error) {
	panic(
		wire.Build(
			controller.NewHealthController,
		),
	)
}
