package ioc

import (
	"go.uber.org/zap"
)

func InitZap() {
	logger, _ := zap.NewProduction()

	// 替换全局 logger
	zap.ReplaceGlobals(logger)
}
