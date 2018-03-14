package main

import(
	"github.com/kataras/iris"
	"os"
)

// App instance of iris
var App = iris.New()
func main() {
	// PORT of the application
	var PORT =  os.Getenv("APP_PORT")

	app := App
	
	app.Run(iris.Addr(":" + PORT))
}