package users

import (
	"github.com/labstack/echo/v4"
)

func (ch CreateUsersHandle) SaveUsersEndPoint(c *echo.Context) {
}

func (ch CreateUsersHandle) GetUsersEndPoint(c *echo.Context) {
}

func (ch CreateUsersHandle) GetIdUsersEndPoint(c *echo.Context) {
}

func (ch CreateUsersHandle) UpdateUsersEndPoint(c *echo.Context) {
}

func (ch CreateUsersHandle) DeleteUsersEndPoint(c *echo.Context) {
}

type CreateUsersHandle struct {
}

func NewCreateHandlerUser() *CreateUsersHandle {
	return nil
}