package main

import (
	"lesson/lesson01"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Parent struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Child struct {
	Parent
	Age int
}

func InitDB(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:000000@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(dst...)

	return db
}

func main() {
	db, err := gorm.Open(mysql.Open("root:000000@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	lesson01.Run(db)
	// lesson02.Run(db)
	// lesson03.Run(db)
	// lesson03_02.Run(db)
	// lesson03_03.Run(db)
	// lesson03_04.Run(db)
	// lesson04.Run(db)

	//InitDB(&Parent{}, &Child{})
}
