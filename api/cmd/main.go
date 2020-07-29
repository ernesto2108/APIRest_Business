package main

import (
	"github.com/ernesto2108/APIRest_Business/api/server"
	"github.com/ernesto2108/APIRest_Business/internal/http"
	business "github.com/ernesto2108/APIRest_Business/internal/pkg/business/web"
	categories "github.com/ernesto2108/APIRest_Business/internal/pkg/categories/web"
	menu "github.com/ernesto2108/APIRest_Business/internal/pkg/menu/web"
	products "github.com/ernesto2108/APIRest_Business/internal/pkg/products/web"
	user "github.com/ernesto2108/APIRest_Business/internal/pkg/users/web"
)

func main() {

	//	All Handlers
	UserHandler := user.NewCreateHandlerUser()
	ProductHandler := products.NewCreateProductHandler()
	CategoryHandler := categories.NewCreateCategoryHandler()
	BusinessHandler := business.NewCreateBusinessHandler()
	menuHandler := menu.NewCreateMenuHandler()

	r := http.AllRoutes(UserHandler, ProductHandler, CategoryHandler,
		BusinessHandler, menuHandler)

	server := server.NewServer()
	server.Run(r)
}

