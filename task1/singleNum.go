package main

import "fmt"

var nums = [7]int{1, 2, 3, 4, 2, 3, 4}

func singleNumber(nums []int) {
	freq := make(map[int]int)

	// 统计每个数字出现的次数
	for _, num := range nums {
		freq[num]++
	}
	fmt.Println(freq)

	// 找出只出现一次的数字
	for num, count := range freq {
		if count == 1 {
			fmt.Println(num)
		}
	}
}

func main() {
	singleNumber(nums[:])
}
