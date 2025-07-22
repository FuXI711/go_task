package lesson01

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     uint    `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func Run() {
	db, err := sqlx.Connect("mysql", "root:000000@tcp(127.0.0.1:3306)/gorm?parseTime=true")
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	defer db.Close()
	//查询技术部员工信息
	var books []Book

	err = db.Select(&books, "SELECT * FROM books where price > ?", 50)

	fmt.Println(books)
}
