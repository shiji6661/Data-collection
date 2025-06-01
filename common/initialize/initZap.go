package initialize

import (
	"fmt"
	"go.uber.org/zap"
)

func ZapInit() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{
		"./Zap/zap.log",
		"stdout",
	}
	logger, err := config.Build()
	if err != nil {
		return
	}

	zap.ReplaceGlobals(logger)
	fmt.Println("zap init success")
}
