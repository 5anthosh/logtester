package main

import (
	"errors"
	"time"

	"go.uber.org/zap"
)

var errValue = errors.New("dummy error")

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second * 2)
		logger.Error("this is error logging", zap.Error(errValue))
		time.Sleep(time.Second * 2)
		logger.Info("this is info logging", zap.String("VALUE", "string value"))
	}
}
