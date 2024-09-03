package main

import (
	_ "github.com/KozlovNikolai/marketplace/docs"
	"github.com/KozlovNikolai/marketplace/internal/app/transport/httpserver"
	"github.com/KozlovNikolai/marketplace/internal/pkg/config"
)

// @title 	Shop Service API
// @version	1.0
// @description An Shop service API in Go using Gin framework
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host 	localhost:8443
// @BasePath /
func main() {
	config.MustLoad()
	server := httpserver.NewServer()

	server.Run()
}
