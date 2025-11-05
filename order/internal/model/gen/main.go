package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shop/order/internal/model"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/order?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.OrderInfo{}, &model.ShoppingCart{}, &model.OrderGoods{})
	if err != nil {
		return
	}
}
