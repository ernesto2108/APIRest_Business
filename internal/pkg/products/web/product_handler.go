package products

import (
	"github.com/labstack/echo/v4"
)

func (ch CreateProductHandler) SaveProductEndPoint(c echo.Context) {
}

func (ch CreateProductHandler) GetProductEndPoint(c echo.Context) {
}

func (ch CreateProductHandler) GetIdProductEndPoint(c echo.Context) {
}

func (ch CreateProductHandler) UpdateProductEndPoint(c echo.Context) {
}

func (ch CreateProductHandler) DeleteProductEndPoint(c echo.Context) {
}

type CreateProductHandler struct {
}

func NewCreateProductHandler() *CreateProductHandler {
	return nil
}
