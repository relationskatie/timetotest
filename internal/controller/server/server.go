package server

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/relationskatie/timetotest/config"
	"github.com/relationskatie/timetotest/internal/controller"
	storage2 "github.com/relationskatie/timetotest/internal/storage"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"testing"
)

var _ controller.Controller = (*Controller)(nil)

type Controller struct {
	log      *zap.Logger
	cfg      *config.Config
	pool     *pgxpool.Pool
	storage2 storage2.Interface
	srv      *echo.Echo
}

func New(log *zap.Logger, cfg *config.Config, pool *pgxpool.Pool, storage2 storage2.Interface) (*Controller, error) {
	ctrl := &Controller{
		log:      log,
		cfg:      cfg,
		pool:     pool,
		storage2: storage2,
		srv:      echo.New(),
	}
	ctrl.configureMiddlewares()
	ctrl.configureRoutes()
	return ctrl, nil
}

func (ctrl *Controller) configureRoutes() {
	api := ctrl.srv.Group("/api")
	{
		api.POST("/add_user/", ctrl.handleAddNewUser)
		api.PATCH("/change_user/", ctrl.handleChangeUser)
		api.GET("/return_all_users/", ctrl.handleGetAllUsers)
		api.DELETE("/delete_user/:username", ctrl.handleDeleteUser)
		api.GET("/user/:id", ctrl.handleGetUserByID)
	}
}

func (ctrl *Controller) configureMiddlewares() {
	ctrl.srv.Use(middleware.Logger())
}

func (ctrl *Controller) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		err := ctrl.srv.Start(ctrl.cfg.Server.GetBindAddress())
		ctrl.log.Info("server start")
		if err != nil {
			ctrl.log.Error("failed to start server", zap.Error(err))
			cancel()
		}
	}()
	return ctx.Err()
}
func (ctrl *Controller) Shutdown(ctx context.Context) error {
	return ctrl.srv.Shutdown(ctx)
}
func testController(t testing.TB, store storage2.Interface) *Controller {
	t.Helper()
	ctrl := &Controller{
		log:      zaptest.NewLogger(t),
		storage2: store,
		srv:      echo.New(),
	}
	return ctrl
}
