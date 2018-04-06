package main

import (
	db "WhereIsMyDriver/databases"
	"WhereIsMyDriver/routers"
	"os"

	"github.com/kataras/iris"
)

func main() {
	// PORT of the application
	var PORT = os.Getenv("PORT")
	app := routers.IrisApp()

	// Make run migration
	db.RunMigration()
	db.SeedData()
	app.Run(iris.Addr(":" + PORT))
}
