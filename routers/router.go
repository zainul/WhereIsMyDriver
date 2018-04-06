package routers

import (
	"WhereIsMyDriver/controllers"

	"github.com/kataras/iris"
)

// IrisApp iris App is router
func IrisApp() *iris.Application {
	app := iris.New()
	app.Get("/drivers", controllers.GetDrivers)
	app.Put("/drivers/:id/location", controllers.UpdateLocation)
	app.OnErrorCode(iris.StatusNotFound)
	return app
}
