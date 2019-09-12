package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"log"
	"math/rand"
	"os"
	"time"
)

//最大号码
const rateMax  =	16

type gift struct {
	id 		int		// 奖品ID
	name 	string	// 奖品名称
	pic 	string
	link 	string
	inuse	bool
	rate 	int 	// 中奖概率，十分之N,0-9
	rateMin	int		// 大于等于，中奖的最小号码,0-10
	rateMax	int		// 小于，中奖的最大号码,0-10
}

var logger *log.Logger

func newGifts() *[5]gift {
	giftList := new([5]gift)
	gift1 := gift{
		id:      0,
		name:    "富强福",
		pic:     "富强福.jpg",
		link:    "富强福.html",
		inuse:   true,
		rateMin: 1,
		rateMax: 4,
	}
	giftList[0] = gift1
	gift2 := gift{
		id:      1,
		name:    "和谐福",
		pic:     "和谐福.jpg",
		link:    "和谐福.html",
		inuse:   true,
		rateMin: 5,
		rateMax: 8,
	}
	giftList[1] = gift2
	gift3 := gift{
		id:      2,
		name:    "爱国福",
		pic:     "爱国福.jpg",
		link:    "爱国福.html",
		inuse:   true,
		rateMin: 9,
		rateMax: 12,
	}
	giftList[2] = gift3
	gift4 := gift{
		id:      3,
		name:    "友善福",
		pic:     "友善福.jpg",
		link:    "友善福.html",
		inuse:   true,
		rateMin: 13,
		rateMax: 16,
	}
	giftList[3] = gift4
	gift5 := gift{
		id:      4,
		name:    "敬业福",
		pic:     "敬业福.jpg",
		link:    "敬业福.html",
		inuse:   true,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[4] = gift5

	return giftList
}

func InitLog()  {
	f, _ := os.Create("/Users/zhaoliangchen/Documents/demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

type LotteryController struct {
	Ctx 	iris.Context
}

func (c* LotteryController) Get() string {
	giftList := newGifts()
	return fmt.Sprintf("%v\n", giftList)
}

func (c* LotteryController) GetLucky() map[string]interface{}  {
	//uid , _ := c.Ctx.URLParamInt("uid")
	//rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	code := luckyCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = ok
	giftList := newGifts()
	for _, data := range giftList{
		if !data.inuse {
			continue
		}
		if data.rateMin <= int(code) && int(code) <= data.rateMax {
			ok = true
			fmt.Printf("抽中了%s\n",data.name)
			result["name"] = data.name
			result["success"] = ok
			break
		}
	}
	return result
}

func luckyCode() int32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(rateMax))
}