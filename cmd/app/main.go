package main

import (
	"github.com/relationskatie/timetotest/config"
	"github.com/relationskatie/timetotest/internal/controller"
	"github.com/relationskatie/timetotest/internal/controller/server"
	pkg "github.com/relationskatie/timetotest/internal/pkg/pgx"
	storage2 "github.com/relationskatie/timetotest/internal/storage"
	"github.com/relationskatie/timetotest/internal/storage/pgx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	createAppWithFx()
}

func createAppWithFx() {
	app := fx.New(
		fx.Provide(
			zap.NewProduction,
			config.New,
			pkg.New,
			pgx.New,
			fx.Annotate(server.New,
				fx.As(new(controller.Controller))),
			fx.Annotate(pgx.New, fx.As(new(storage2.Interface))),
		),
		fx.Invoke(
			controller.StartFx,
		),
	)
	app.Run()
}
