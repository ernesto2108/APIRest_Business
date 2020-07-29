package http

import (
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"net/http"

	//	services api RestFull
	business "github.com/ernesto2108/APIRest_Business/internal/pkg/business/web"
	categories "github.com/ernesto2108/APIRest_Business/internal/pkg/categories/web"
	menu "github.com/ernesto2108/APIRest_Business/internal/pkg/menu/web"
	products "github.com/ernesto2108/APIRest_Business/internal/pkg/products/web"
	user "github.com/ernesto2108/APIRest_Business/internal/pkg/users/web"
)

func AllRoutes(

	//	All Routes
	userHandler *user.CreateUsersHandle,
	productHandler *products.CreateProductHandler,
	categoryHandler *categories.CreateCategoryHandler,
	businessHandler *business.CreateBusinessHandler,
	menuHandler *menu.CreateMenuHandler,

	) *echo.Echo {

	r := echo.New()

	r.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	/*	Global Middleware  */
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	//	consultas admin
	U1 := r.Group("admin")
	U1.POST("/users", userHandler.SaveUsersEndPoint)

	//	consultas user
	U2 := r.Group("user")
	U2.GET("/users", userHandler.GetUsersEndPoint)

	return r
}
