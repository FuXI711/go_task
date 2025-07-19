package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串作为基准
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		// 比较当前字符串与prefix的公共前缀
		j := 0
		fmt.Println("j	", j)
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		fmt.Println("j", j)
		// 更新prefix为当前找到的公共前缀
		prefix = prefix[:j]

		// 如果prefix已经为空，可以提前结束
		if prefix == "" {
			break
		}
	}

	return prefix
}

func main() {
	strs := []string{"flight", "flow", "flower"}
	fmt.Println(longestCommonPrefix(strs)) // 输出 "fl"
}
