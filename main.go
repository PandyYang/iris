package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris/controller"
)

func main() {

	app := iris.New()
	app.RegisterView(iris.HTML("./view", ".html"))
	//注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controller.MovieController))
	app.Run(
		iris.Addr("localhost:8088"),
	)
}



