package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
)

type LotteryController struct {
	Ctx iris.Context
}

func (c *LotteryController) Get() string {
	b, _ := json.Marshal(gifts)
	return string(b)
}

const (
	giftTypeCoin      = iota // 虚拟币
	giftTypeCoupon           // 优惠券，不相同的编码
	giftTypeCouponFix        // 优惠券，相同的编码
	giftTypeRealSmall        // 实物小奖
	giftTypeRealLarge        // 实物大奖
)

// 奖品信息
type gift struct {
	id       int      // 奖品ID
	name     string   // 奖品名称
	gtype    int      // 奖品类型
	data     string   // 奖品的数据（特定的配置信息，如：虚拟币面值，固定优惠券的编码）
	datalist []string // 奖品数据集合（特定的配置信息，如：不同的优惠券的编码）
	total    int      // 总数，0 不限量
	left     int      // 剩余数
	inuse    bool     // 是否使用中
	rateMin  int      // 大于等于，中奖的最小号码,0-10000
	rateMax  int      // 小于，中奖的最大号码,0-10000
}

var gifts []*gift
func InitGifts() {
	rateStart := 1
	names := [...]string{"虚拟币", "优惠券", "实物小奖", "实物大奖"}
	for i := 0; i < len(names); i++ {
		gift := gift{
			id:       i,
			name:     names[i],
			gtype:    0,
			data:     "",
			datalist: nil,
			total:    0,
			left:     0,
			inuse:    false,
			rateMin:  rateStart,
			rateMax:  rateStart + 100,
		}
		rateStart = gift.rateMax + 1
		if i == len(names)-1 {
			gift.rateMin = 0
			gift.rateMax = 0
		}
		gifts = append(gifts, gift)
	}
	fmt.Printf("%v", gifts)
}