package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"lottery-golang/微博抢红包/controller"
)

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&controller.LotteryController{})
	go controller.FetchPackageMoney()
	return app
}

func main()  {
	app := newApp()
	app.Run(iris.Addr("localhost:8080"))
}