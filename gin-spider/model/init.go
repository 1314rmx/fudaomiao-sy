package model

import (
	"log"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB = Init()

func Init() *gorm.DB {
	dsn := "fudaomiao:fudaomiao123@tcp(10.99.1.40:3306)/fudaomiao?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Println("连接数据库失败", err)
		return nil
	}
	return db
}
