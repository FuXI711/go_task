package main

import (
	"fmt"
)

func addTen(num *int) {
	*num = *num * 2
}

func main() {
	value := []int{1, 2, 3}

	for i := range value {
		addTen(&value[i])
	}

	fmt.Println("修改后的值:", value)
}
