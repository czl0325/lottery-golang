package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"lottery-golang/支付宝五福/controller"
)

func newApp() *iris.Application  {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&controller.LotteryController{})
	controller.InitLog()
	return app
}

func main()  {
	app := newApp()
	app.Run(iris.Addr("localhost:8080"))
}
