package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/ivanovishado/proxy-app/api/middleware"
	"github.com/kataras/iris"
)

// RedirectionHandler should redirect traffic
func RedirectionHandler(app *iris.Application) {
	app.Get("/ping", middleware.ProxyMiddleware, proxyHandler)
}

func proxyHandler(c iris.Context) {
	fmt.Println(middleware.AppRequests)
	response, err := json.Marshal(middleware.AppRequests)
	if err != nil {
		c.JSON(iris.Map{"status": 400, "result": "parse error"})
		return
	}
	c.JSON(iris.Map{"result": string(response)})
}
