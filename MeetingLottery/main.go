package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"lottery-golang/MeetingLottery/controller"
)

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&controller.LotteryController{})
	return app
}

func main()  {
	app := newApp()
	app.Run(iris.Addr("localhost:8080"))
}
