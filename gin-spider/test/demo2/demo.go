package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
	Age  int
}

func main() {
	/*
		dsn := "root:520769rmx@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
		db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		sqlDB, _ := db.DB()
		//SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)
		var u1 User
		db.First(&u1)
		fmt.Println(u1)
	*/

	//key := "d0m6rQJBKdw2yub7"
	//pwd := "163155"
	//cmd := exec.Command("node", "js/encryptPWD.js", pwd, key)
	//out, err := cmd.CombinedOutput()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(out))

	//c := model.Initcolly("202330325", "163155")
	//fmt.Println(c)

	timestamp := time.Now().UnixNano() / 1e6
	timestampStr := fmt.Sprintf("%d", timestamp)
	fmt.Println(timestampStr)

}
