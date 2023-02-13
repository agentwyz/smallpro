// package main

// import (
// 	"fmt"
// 	_ "fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// /*
// 数据表<--->结构体
// 数据行<--->结构体实现
// 字段<--->结构体字段
// */

// type UserInfo struct {
// 	ID uint
// 	Name string
// 	Gender string
// 	Hobby string
// }


// func main() {
// 	dsn := "root:123456@tcp(localhost)/db01?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic(err)
// 	}
	
// 	db.AutoMigrate(&UserInfo{})

// 	//创建数据行, 加上一个Users
// 	//u1 := UserInfo{1, "wyz", "mam", "basketball"}
// 	//db.Create(&u1)

// 	//查询
// 	var u UserInfo
// 	db.First(&u) //查询表中第一条数据, 保存到
// 	fmt.Printf("u: %#v\n", u)
	
// 	//更新
// 	db.Model(&u).Update("hobby", "篮球")

// 	//删除操作
// 	db.Delete(&u)
	
// 	//
// }