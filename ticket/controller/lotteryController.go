package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"math/rand"
	"time"
)

type LotteryController struct {
	Ctx iris.Context
}

func (c* LotteryController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Intn(10)
	var prize string
	switch code {
	case 1:
		prize = "一等奖"
	case 2,3:
		prize = "二等奖"
	case 4,5,6:
		prize = "三等奖"
	default:
		return fmt.Sprintf("尾号为1是一等奖<br/>" +
			"尾号为2,3是二等奖<br/>" +
			"尾号为4，5，6是三等奖<br/>" +
			"其余不中奖<br/>" +
			"您没有中奖!")
	}
	return fmt.Sprintf("尾号为1是一等奖<br/>" +
		"尾号为2,3是二等奖<br/>" +
		"尾号为4，5，6是三等奖<br/>" +
		"其余不中奖<br/>" +
		"您的尾号是%d,恭喜您获得%s",code,prize)
}
