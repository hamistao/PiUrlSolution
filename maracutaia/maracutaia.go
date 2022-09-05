package main

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
)

func consume(words chan string) {
	for word := range words {
		n, _ := strconv.Atoi(word)
		if isPalindrome(word) && isPrime(int64(n)) {
			url := fmt.Sprintf("http://%v.com/", word)
			r, _ := http.Get(url)
			if r != nil {
				fmt.Println(url)
				fmt.Println("foi")
			}
		}
	}
}

func join(alg [5]int) string {
	out := ""
	for i := 0; i < len(alg); i++ {
		out += strconv.Itoa(alg[i])

	}
	for i := len(alg) - 2; i >= 0; i-- {
		out += strconv.Itoa(alg[i])
	}
	return out
}

func isPrime(n int64) bool { // function to check if number is prime
	return big.NewInt(n).ProbablyPrime(0)
}

func isPalindrome(s string) bool { // function to check if word is palindrome
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	words := make(chan string, 100)
	for i := 0; i < 3; i++ {
		go consume(words)
	}
	current := 4
	ind := 4
	alg := [5]int{0, 0, 0, 0, 0}
	for ind >= 0 {
		words <- join(alg)
		alg[current]++
		if alg[current] == 10 {
			for current >= 0 && alg[current] == 10 {
				alg[current] = 0
				current--
				if current >= 0 {
					alg[current]++
				}
			}
			if current < ind {
				ind--
			}
			current = 4
		}
	}
}
