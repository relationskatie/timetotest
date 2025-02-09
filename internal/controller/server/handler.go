package server

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/relationskatie/timetotest/internal/modles"
	"net/http"
)

func (ctrl *Controller) handleAddNewUser(ctx echo.Context) error {
	var req modles.AddUserRequest
	if err := ctx.Bind(&req); err != nil {
		ctrl.log.Error("failed to bind add new user request")
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	data := modles.UserDTO{
		ID:        uuid.New(),
		Name:      req.Name,
		Username:  req.Username,
		Age:       req.Age,
		Telephone: req.Telephone,
	}
	err := ctrl.storage2.User().AddNewUser(ctx.Request().Context(), data)
	if err != nil {
		ctrl.log.Error("failed to add new user")
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, nil)
}

func (ctrl *Controller) handleChangeUser(ctx echo.Context) error {
	var req modles.ChangeUserRequest
	if err := ctx.Bind(&req); err != nil {
		ctrl.log.Error("failed to bind change user request")
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	data := modles.ChangeUserDTO{
		Name:      req.Name,
		Age:       req.Age,
		Telephone: req.Telephone,
		Username:  req.Username,
	}
	err := ctrl.storage2.User().ChangeUser(ctx.Request().Context(), data)
	if err != nil {
		ctrl.log.Error("failed to change user")
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, struct{}{})
}

func (ctrl *Controller) handleDeleteUser(ctx echo.Context) error {
	name := ctx.Param("name")
	if len(name) == 0 {
		ctrl.log.Error("failed to delete user")
		return ctx.JSON(http.StatusBadRequest, "name is required")
	}
	err := ctrl.storage2.User().DeleteUserByUsername(ctx.Request().Context(), name)
	if err != nil {
		ctrl.log.Error("failed to delete user")
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (ctrl *Controller) handleGetAllUsers(ctx echo.Context) error {
	data, err := ctrl.storage2.User().GetAllUsers(ctx.Request().Context())
	if err != nil {
		ctrl.log.Error("failed to get all users")
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, data)
}

func (ctrl *Controller) handleGetUserByID(ctx echo.Context) error {
	id := ctx.Param("id")
	ID, err := uuid.Parse(id)
	if err != nil {
		ctrl.log.Error("failed to parse user id")
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid UUID format"})
	}

	data, err := ctrl.storage2.User().GetUserByID(ctx.Request().Context(), ID)
	if err != nil {
		ctrl.log.Error("failed to get user by id")
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, data)
}
