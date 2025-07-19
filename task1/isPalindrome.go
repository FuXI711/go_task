package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		fmt.Println("reversed", reversed)
		x /= 10
		fmt.Println("x", x)
	}
	return x == reversed || x == reversed/10
}

func main() {
	fmt.Println(isPalindrome(121121)) // true
	// fmt.Println(isPalindrome(-121))  // false
	// fmt.Println(isPalindrome(10))    // false
	// fmt.Println(isPalindrome(12321)) // true
}
