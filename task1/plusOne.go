package main

import "fmt"

func plusOne(digits []int) []int {
	carry := 1                              // 初始加1
	for i := len(digits) - 1; i >= 0; i-- { //digits[i] = 3   2   1
		sum := digits[i] + carry //3+1
		digits[i] = sum % 10     //0
		carry = sum / 10         //4
		if carry == 0 {
			break // 如果没有进位，可以提前结束
		}
	}
	// 如果最后仍有进位，需要在数组最前面插入1
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits)) // 输出 [1 2 4]
}
