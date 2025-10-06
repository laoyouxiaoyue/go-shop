package main

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shop/goods/model"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goods?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Panicf("数据库连接失败 %v", zap.Error(err))
		return
	}
	err = db.AutoMigrate(&model.Category{}, &model.Brands{},
		&model.GoodsCategoryBrand{}, &model.Brands{}, &model.Goods{})
	if err != nil {
		return
	}
}
