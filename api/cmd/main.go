package main

import (
	"github.com/ernesto2108/APIRest_Business/api/server"
	"github.com/ernesto2108/APIRest_Business/internal/http"
)

func main() {
	r := http.AllRoutes()
	server := server.NewServer()
	server.Run(r)
}
