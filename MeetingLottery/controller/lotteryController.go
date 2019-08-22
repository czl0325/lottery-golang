package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type LotteryController struct {
	Ctx  iris.Context
}

var userList []string
var mu sync.Mutex

func (c* LotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("当前参与抽奖的人数为 %d\n", count)
}

// POST http://localhost:8080/import
// 导入抽奖人员，字段users="user1,user2,user3"，逗号隔开
func (c *LotteryController) PostImport() string  {
	mu.Lock()
	defer mu.Unlock()

	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")
	count1 := len(userList)
	for _, data := range users {
		data = strings.TrimSpace(data)
		if len(data) > 0 {
			userList = append(userList, data)
		}
	}
	count2 := len(userList)
	return fmt.Sprintf("当前参与抽奖总人数为 %d, 成功导入 %d 人",count2, count2-count1)
}

// GET http://localhost:8080/lucky
// 获取中奖人员
func (c *LotteryController) GetLucky() string {
	count := len(userList)
	if count > 1 {
		seed := time.Now().UnixNano()
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
		user := userList[index]
		userList = append(userList[0:index], userList[index+1:]...)
		return fmt.Sprintf("当前中奖用户: %s, 剩余用户数: %d\n", user, len(userList))
	} else if count == 1 {
		user := userList[0]
		userList = userList[0:0]
		return fmt.Sprintf("当前中奖用户: %s, 剩余用户数: %d\n", user, len(userList))
	} else {
		return fmt.Sprintf("当前无用户参与抽奖")
	}
}