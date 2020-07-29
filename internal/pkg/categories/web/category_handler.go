package categories

import (
	"github.com/labstack/echo/v4"
)

func (ch CreateCategoryHandler) SaveCategoryEndPoint(c echo.Context) {
}

func (ch CreateCategoryHandler) GetCategoryEndPoint(c echo.Context) {
}

func (ch CreateCategoryHandler) GetIdCategoryEndPoint(c echo.Context) {
}

func (ch CreateCategoryHandler) UpdateCategoryEndPoint(c echo.Context) {
}

func (ch CreateCategoryHandler) DeleteCategoryEndPoint(c echo.Context) {
}

type CreateCategoryHandler struct {
}

func NewCreateCategoryHandler() *CreateCategoryHandler {
	return nil
}
