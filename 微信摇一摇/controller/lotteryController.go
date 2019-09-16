package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"math/rand"
	"sync"
	"time"
)

type LotteryController struct {
	Ctx iris.Context
}

func (c *LotteryController) Get() string {
	b, _ := json.Marshal(gifts)
	return string(b)
}

var gifts []*gift
var mu sync.Mutex
func (c* LotteryController) GetLucky() string {
	mu.Lock()
	defer mu.Unlock()

	index := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(rateMax))
	for _, data := range gifts {
		if index >= int32(data.rateMin) && index <= int32(data.rateMax) {
			fmt.Printf("%v\n",data)
			if data.left <= 0 {
				fmt.Printf("奖品%s没有库存了\n",data.name)
				return fmt.Sprintf("奖品%s没有库存了",data.name)
			} else {
				data.left = data.left - 1
				fmt.Printf("抽中奖品%s库存还剩下%d\n",data.name,data.left)
				return fmt.Sprintf("抽中奖品%s库存还剩下%d",data.name,data.left)
			}
		}
	}
	fmt.Printf("当前数字为%d,没有中奖\n",index)
	return fmt.Sprintf("当前数字为%d,没有中奖",index)
}

const (
	giftTypeCoin      = iota // 虚拟币
	giftTypeCoupon           // 优惠券，不相同的编码
	giftTypeCouponFix        // 优惠券，相同的编码
	giftTypeRealSmall        // 实物小奖
	giftTypeRealLarge        // 实物大奖
)
var rateMax int

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
			total:    100,
			left:     100,
			inuse:    true,
			rateMin:  rateStart,
			rateMax:  rateStart + 100,
		}
		rateStart = gift.rateMax + 1
		rateMax = gift.rateMax
		if i == len(names)-1 {
			gift.rateMin = 0
			gift.rateMax = 0
		}
		gifts = append(gifts, &gift)
	}
	fmt.Printf("%v", &gifts)
}