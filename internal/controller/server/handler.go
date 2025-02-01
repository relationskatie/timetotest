package server

import (
	"github.com/labstack/echo/v4"
	"github.com/relationskatie/timetotest/internal/modles"
	"net/http"
)

func (ctrl *Controller) handleAddNewUser(ctx echo.Context) error {
	var req modles.AddUserRequest
	if err := ctx.Bind(&req); err != nil {
		ctrl.log.Error("failed to bind add new user request")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//data, err := ctrl.storage2
	return echo.NewHTTPError(http.StatusCreated)
}

func (ctrl *Controller) handleChangeUser(ctx echo.Context) error {
	//need add logic to add in db
	return echo.NewHTTPError(http.StatusNoContent)
}

func (ctrl *Controller) handleDeleteUser(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNoContent)
}

func (ctrl *Controller) handleGetAllUsers(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNoContent)
}
