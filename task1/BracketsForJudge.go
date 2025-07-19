package main

import "fmt"

func IsValidRelaxed(s string) bool {
	count := make(map[byte]int)
	pairs := map[byte]byte{')': '(', '}': '{', ']': '['}

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(', '{', '[':
			count[c]++
		case ')', '}', ']':
			count[pairs[c]]--
			if count[pairs[c]] < 0 {
				return false
			}
		default:
			return false
		}
	}

	// 检查所有括号数量是否平衡
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsValidRelaxed("()[]{}"))
}
