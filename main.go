package main

import (
	handlers "github.com/ivanovishado/proxy-app/api/handlers"
	"github.com/ivanovishado/proxy-app/api/middleware"
	server "github.com/ivanovishado/proxy-app/api/server"
	utils "github.com/ivanovishado/proxy-app/api/utils"
)

func main() {
	utils.LoadEnvVars()
	middleware.SetPriorityLevels()
	app := server.SetUp()
	handlers.RedirectionHandler(app)
	server.RunServer(app)
}
