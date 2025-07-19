package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		fmt.Println("hashMap", hashMap)
		if j, ok := hashMap[complement]; ok {
			fmt.Println("j", []int{j, i})
			return []int{j, i}
		}

		hashMap[num] = i
	}

	return nil
}

func main() {
	intervals5 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("输出5:", twoSum(intervals5, 11))
}
