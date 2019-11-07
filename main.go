package main

import (
	"github.com/yangxinwei/gin-demo-api/controllers"
	"github.com/yangxinwei/gin-demo-api/bootstrap"
	_ "github.com/yangxinwei/gin-demo-api/docs"
	"github.com/yangxinwei/gin-demo-api/routers"
)

// @title API
// @version 0.0.1
// @description

// @contact.name API Support
// @contact.email xxx@xxx.com
// @host
// @BasePath /api/v1/

func main() {
	app := bootstrap.APP{}

	app.Controller = new(controllers.BaseController)
	app.Initialize()
	app.Router = routers.SetupRouter(app.Controller)

	app.Run(":8080")
}
