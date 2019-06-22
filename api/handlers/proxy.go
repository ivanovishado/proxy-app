package handlers

import (
	"github.com/ivanovishado/proxy-app/api/middlewares"
	"github.com/kataras/iris"
)

// RedirectionHandler should redirect traffic
func RedirectionHandler(app *iris.Application) {
	app.Get("/ping", middlewares.ProxyMiddleware, proxyHandler)
}

func proxyHandler(c iris.Context) {
	c.JSON(iris.Map{"result": "ok"})
}
