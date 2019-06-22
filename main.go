package main

import (
	handlers "github.com/ivanovishado/proxy-app/api/handlers"
	middleware "github.com/ivanovishado/proxy-app/api/middlewares"
	server "github.com/ivanovishado/proxy-app/api/server"
	utils "github.com/ivanovishado/proxy-app/api/utils"
)

func main() {
	utils.LoadEnvVars()
	app := server.SetUp()
	middleware.InitQueue()
	handlers.RedirectionHandler(app)
	server.RunServer(app)
}
