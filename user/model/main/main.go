package main

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"shop/user/model"
	"time"
)

func genMd5(code string) (string, error) {
	md5Ctx := md5.New()
	_, err := md5Ctx.Write([]byte(code))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5Ctx.Sum(nil)), nil
}
func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/shop?charset=utf8&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&model.User{})

}
