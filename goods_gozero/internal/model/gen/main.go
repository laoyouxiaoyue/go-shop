package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shop/goods_gozero/internal/model"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goods?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Brands{}, &model.Goods{}, &model.Category{}, &model.Banner{})
	if err != nil {
		return
	}
}
