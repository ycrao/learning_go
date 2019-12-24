package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MysqlDB *gorm.DB

type User struct {
	// Id   int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	gorm.Model
	Age  int    `gorm:"size:11;DEFAULT NULL" json:"age"`
	Name string `gorm:"size:255;DEFAULT NULL" json:"name"`
}

func main() {
	var err error
	MysqlDB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test_dev?charset=utf8mb4")
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	} else {
		fmt.Println("connect database success")
		MysqlDB.SingularTable(true)
		MysqlDB.AutoMigrate(&User{})
		fmt.Println("create table success")
	}
	defer MysqlDB.Close()
	createUser()
}

func createUser() {
	user := User{
		Name: "张三",
		Age: 12,
	}
	MysqlDB.Create(&user)
}
