package test

import (
	"fmt"
	"keep.com/godemo/cfg"
	"keep.com/godemo/logger"
	"testing"
	"time"
)

func TestCase(t *testing.T) {
	logger.Init()
	cfg.Init()

	//metrics.Init()
	//err := models.InitDB()
	//
	//if err != nil {
	//	logger.Panic("init mysql failed, err: %v", err)
	//}

	//r := api.Init()
	//httpx.Init(r)

	err := testMe()
	if err != nil {
		logger.Errorf("test err: %v", err.Error())
	} else {
		logger.Info("PASS!")
	}

	time.Sleep(time.Minute / 2)
}

func testMe() (err error) {
	return fmt.Errorf("mock err")
}

