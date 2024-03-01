package main

import (
	"github.com/Ferriem/Iris/code/web/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))
	//Controller
	mvc.New(app.Party("/movie")).Handle(new(controllers.MovieController))
	app.Run(
		iris.Addr("localhost:8080"),
	)

}
