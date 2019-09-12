package main

import (
	"fmt"
	"github.com/kataras/iris/httptest"
	"sync"
	"testing"
)

/**
	测试用例一定要
	go get -u github.com/gavv/httpexpect
 */
func Test(t *testing.T)  {
	e := httptest.New(t, newApp())

	var wg sync.WaitGroup

	for i:=0; i<100; i++  {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			e.POST("/import").WithFormField("users", fmt.Sprintf("test_u%d", i)).Expect().Status(httptest.StatusOK)
		}(i)
	}

	wg.Wait()
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前参与抽奖的人数为 100\n")
	e.GET("/lucky").Expect().Status(httptest.StatusOK)
	e.GET("/").Expect().Status(httptest.StatusOK).Body().Equal("当前参与抽奖的人数为 99\n")
}
