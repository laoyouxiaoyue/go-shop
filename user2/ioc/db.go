package ioc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	"shop/user2/config"
)

func InitDB() *gorm.DB {
	fmt.Println(config.Cf.Mysql.Dsn)
	db, err := gorm.Open(mysql.Open(config.Cf.Mysql.Dsn), &gorm.Config{})
	slog.Info("数据库Dsn:", config.Cf.Mysql.Dsn)
	if err != nil {
		panic("数据库连接失败")
	}
	slog.Info(fmt.Sprintf("数据库连接成功"))
	return db
}
