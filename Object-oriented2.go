package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Println("姓名为", e.Name, "年龄为", e.Age, "员工ID为", e.EmployeeID)
}

func main() {
	person := Person{Name: "张三", Age: 20}
	employee := Employee{Person: person, EmployeeID: 1001}
	employee.PrintInfo()
}
