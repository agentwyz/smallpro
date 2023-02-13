package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//at a model
type UserInfo struct {
	gorm.Model   //如果一个结构体中含有一个ID字段, 那么将会默认作为主键
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"` //给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`          //这玩意就相当于一个注解, 表示忽略这个字段
}

//注意自定义表的名称
//首先默认表的名称是结构体的复数

//可以先创建一个ID作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64	
}

func (Animal) TableName() string {
	return "mouse"
}

func main() {
	//gorm.DefaultTableNameHandler
	// gorm.DefaultTableNameHanlder = func(db *gorm.DB,defaultTableName string){
	// //设置一个前缀 	
	//return "sys_" + defaultTableName
	// }
	
	dsn := "root:123456@tcp(localhost)/db01?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//db.SkipDefaultTransaction()
	
	//指定一个名称
	//db.Table("wyz").AutoMigrate(&Animal{})
	
	//自动迁移
	db.AutoMigrate(&Animal{})
}
