package main

import (
	"github.com/relationskatie/timetotest/config"
	"github.com/relationskatie/timetotest/internal/controller"
	"github.com/relationskatie/timetotest/internal/controller/server"
	pkg "github.com/relationskatie/timetotest/internal/pkg/pgx"
	storage2 "github.com/relationskatie/timetotest/internal/storage"
	"github.com/relationskatie/timetotest/internal/storage/pgx"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(createAppWith()).Run()
}

func createAppWith() fx.Option {
	return fx.Options(
		fx.Provide(
			zap.NewProduction,
			config.New,
			pkg.New,
			pgx.New,

			fx.Annotate(server.New,
				fx.As(new(controller.Controller))),

			fx.Annotate(pgx.New,
				fx.As(new(storage2.Interface)))),
		fx.Invoke(
			controller.StartFx,
		),

		fx.WithLogger(loggerForFX),
	)
}

func loggerForFX(log *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: log}
}
