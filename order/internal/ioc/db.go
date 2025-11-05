package ioc

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shop/order/internal/config"
	"time"
)

func InitDB(c config.Config) *gorm.DB {
	startTime := time.Now()
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		zap.S().Panicf("数据库连接失败 %v", zap.Error(err))
		return nil
	}
	connectionTime := time.Since(startTime)
	zap.L().Info("数据库连接成功", zap.String("DSN", c.Mysql.Dsn), zap.Duration("cost", connectionTime))
	return db
}
