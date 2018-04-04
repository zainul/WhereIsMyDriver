package main

import (
	"WhereIsMyDriver/controllers"
	db "WhereIsMyDriver/databases"
	"os"

	"github.com/kataras/iris"
)

func irisApp() *iris.Application {
	app := iris.New()
	app.Get("/drivers", controllers.GetDrivers)
	app.Put("/drivers/:id/location", controllers.UpdateLocation)
	app.OnErrorCode(iris.StatusNotFound)
	return app
}

func main() {
	// PORT of the application
	var PORT = os.Getenv("PORT")
	app := irisApp()

	// Make run migration
	db.RunMigration()
	db.SeedData()
	app.Run(iris.Addr(":" + PORT))
}
