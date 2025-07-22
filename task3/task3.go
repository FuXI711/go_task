package lesson01

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Employee 结构体对应employees表
type Employee struct {
	ID         uint    `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func Run() {
	db, err := sqlx.Connect("mysql", "root:000000@tcp(127.0.0.1:3306)/gorm?parseTime=true")
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	defer db.Close()
	//查询技术部员工信息
	var employees []Employee

	err = db.Select(&employees, "SELECT * FROM employees where department = ?", "技术部")

	//  查询工资最高的员工
	var highestPaid Employee
	err = db.Get(&highestPaid, `
		SELECT id, name, department, salary 
		FROM employees 
		ORDER BY salary DESC 
		LIMIT 1
	`)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	fmt.Printf("工资最高的员工信息:\nID: %d\n姓名: %s\n部门: %s\n工资: %.2f\n",
		highestPaid.ID, highestPaid.Name, highestPaid.Department, highestPaid.Salary)
}
