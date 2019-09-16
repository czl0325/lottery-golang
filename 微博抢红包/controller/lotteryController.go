package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"math/rand"
	"sync"
	"time"
)

type LotteryController struct {
	Ctx  iris.Context
}

// 任务结构
type task struct {
	id uint32
	callback chan uint
}

var packageList = new(sync.Map)
var chTask = make(chan task)

// 返回全部红包地址
// GET http://localhost:8080/
func (c* LotteryController) Get() map[uint32][2]int {
	rs := make(map[uint32][2]int)
	packageList.Range(func(key, value interface{}) bool {
		id := key.(uint32)
		list := value.([]uint)
		var money int
		for _, v := range list{
			money += int(v)
		}
		rs[id] = [2]int{len(list), money}
		return true
	})
	return rs
}

// 发红包
// GET http://localhost:8080/set?uid=1&money=100&num=100
func (c *LotteryController) GetSet() string {
	uid, error := c.Ctx.URLParamInt("uid")
	if error != nil {
		return fmt.Sprintf("参数uid异常，%s\n", error)
	}
	money, error := c.Ctx.URLParamFloat64("money")
	if error != nil {
		return fmt.Sprintf("参数money异常，%s\n", error)
	}
	num, error := c.Ctx.URLParamInt("num")
	if error != nil {
		return fmt.Sprintf("参数num异常，%s\n", error)
	}
	moneyTotal := int(money*100)
	if uid < 1 || moneyTotal < num || num < 1 {
		return fmt.Sprintf("参数数值异常，uid=%d, money=%d, num=%d\n", uid, money, num)
	}
	leftMoney := moneyTotal
	leftNum := num
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rMax := 0.55
	if num >= 1000 {
		rMax = 0.01
	} else if num >= 100 {
		rMax = 0.1
	} else if num >= 10 {
		rMax = 0.3
	}
	list := make([]uint, num)
	for leftNum > 0 {
		// 最后一个名额，把剩余的全部给它
		if leftNum == 1 {
			list[num-1] = uint(leftMoney)
			break
		}
		// 剩下的最多只能分配到1分钱时，不用再随机
		if leftMoney == leftNum{
			for i:=num-leftNum; i<num; i++ {
				list[i] = 1
			}
			break
		}
		rMoney := int(float64(leftMoney)*rMax)
		m := r.Intn(rMoney)
		if m < 1 {
			m = 1
		}
		list[num-leftNum] = uint(m)
		leftMoney -= m
		leftNum--
	}
	id := r.Uint32()
	packageList.Store(id, list)
	// 返回抢红包的URL
	fmt.Printf("%v",list)
	return fmt.Sprintf("/get?id=%d&uid=%d&num=%d\n", id, uid, num)
}

func (c* LotteryController) GetGet() string {
	uid, err := c.Ctx.URLParamInt("uid")
	if err != nil {
		return fmt.Sprintf("参数uid异常，%s\n", err)
	}
	id, err := c.Ctx.URLParamInt("id")
	if err != nil {
		return fmt.Sprintf("参数id异常，%s\n", err)
	}
	if uid < 1 || id < 1 {
		return fmt.Sprintf("参数数值异常，uid=%d, id=%d\n", uid, id)
	}
	l, ok := packageList.Load(uint32(id))
	if ok == false {
		return fmt.Sprintf("红包不存在,id=%d\n",id)
	}
	list := l.([]uint)
	if len(list) < 1 {
		return fmt.Sprintf("红包不存在,id=%d\n",id)
	}
	// 更新红包列表中的信息（移除这个金额），构造一个任务
	callback := make(chan uint)
	t := task{id: uint32(id), callback:callback}
	chTask <- t
	money := <- callback
	if money <= 0 {
		fmt.Println(uid, "很遗憾，没能抢到红包")
		return fmt.Sprintf("很遗憾，没能抢到红包\n")
	} else {
		fmt.Println(uid, "抢到一个红包，金额为:", money)
		return fmt.Sprintf("恭喜你抢到一个红包，金额为:%.2f\n", float64(money)/100)
	}
}

// 单线程死循环，专注处理各个红包中金额切片的数据更新（移除指定位置的金额）
func FetchPackageMoney() {
	for {
		t := <- chTask
		// 分配的随机数
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		id := t.id
		l , ok := packageList.Load(id)
		if ok && l!=nil {
			list := l.([]uint)
			// 从红包金额中随机得到一个
			i := r.Intn(len(list))
			money := list[i]
			if len(list) > 1 {
				if i == len(list)-1 {
					packageList.Store(uint32(id), list[:i])
				} else if i == 0 {
					packageList.Store(uint32(id), list[1:])
				} else {
					packageList.Store(uint32(id), append(list[:i], list[i+1:]...))
				}
			} else {
				packageList.Delete(uint32(id))
			}
			t.callback <- money
		} else {
			t.callback <- 0
		}
	}
}