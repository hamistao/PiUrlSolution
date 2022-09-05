package main

import (
	"challenge/pi"
	"fmt"
	"strconv"
)

func isPalindrome(s string) bool { // function to check if word is palindrome
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func cond1() bool {
	fmt.Println("a")
	return false
}

func cond2() bool {
	fmt.Println("b")
	return true
}

func join2(alg [5]int) string {
	out := ""
	for i := 0; i < len(alg); i++ {
		out += strconv.Itoa(alg[i])

	}
	for i := len(alg) - 2; i >= 0; i-- {
		out += strconv.Itoa(alg[i])
	}
	return out
}

func join(alg [5]int) string {
	out := ""
	for i := 0; i < len(alg); i++ {
		out += strconv.Itoa(alg[i])
	}
	return out
}

func main() {
	// fmt.Println(isPalindrome("coioc"))
	// fmt.Println(isPalindrome("coiioc"))
	// fmt.Println(isPalindrome("coiok"))
	// fmt.Println(isPalindrome("koioc"))
	// fmt.Println(isPalindrome("coikoc"))
	// fmt.Println(isPalindrome("cokioc"))
	// fmt.Println(isPalindrome("ckioc"))
	// if cond1() && cond2() {
	// 	fmt.Println("c")
	// }

	// current := 4
	// ind := 4
	// alg := [5]int{1, 2, 0, 4, 0}
	// fmt.Println(join2(alg))

	// for ind >= 0 {
	// 	alg[current]++
	// 	if alg[current] == 10 {
	// 		for current >= 0 && alg[current] == 10 {
	// 			alg[current] = 0
	// 			current--
	// 			if current >= 0 {
	// 				alg[current]++
	// 			}
	// 		}
	// 		if current < ind {
	// 			ind--
	// 		}
	// 		current = 4
	// 	}
	// 	word := join(alg)
	// 	fmt.Println(word)
	// }

	fmt.Println(pi.CalcPi(100))
}
