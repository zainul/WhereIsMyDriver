package main

import (
	db "WhereIsMyDriver/databases"
	"os"

	"github.com/kataras/iris"
)

// App instance of iris
var App = iris.New()

func main() {
	// PORT of the application
	var PORT = os.Getenv("PORT")
	app := App
	// Make run migration
	db.RunMigration()
	app.Run(iris.Addr(":" + PORT))
}
