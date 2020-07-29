package users

import (
	"github.com/labstack/echo/v4"

	"github.com/ernesto2108/APIRest_Business/internal/services/psql"
)

func (ch CreateUsersHandle) SaveUsersEndPoint(c echo.Context) {
}

func (ch CreateUsersHandle) GetUsersEndPoint(c echo.Context) {
}

func (ch CreateUsersHandle) GetIdUsersEndPoint(c echo.Context) {
}

func (ch CreateUsersHandle) UpdateUsersEndPoint(c echo.Context) {
}

func (ch CreateUsersHandle) DeleteUsersEndPoint(c echo.Context) {
}

type CreateUsersHandle struct {
}

func NewCreateHandlerUser(s *services.SqlClient) *CreateUsersHandle {
	return &CreateUsersHandle{}
}