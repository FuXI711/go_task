package main

import (
	"fmt"
)

func addTen(num *int) {
	*num += 10
}

func main() {
	var value int = 5

	addTen(&value)

	fmt.Println("修改后的值:", value)
}
