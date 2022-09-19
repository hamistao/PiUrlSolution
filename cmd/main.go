package main

import (
	"challenge/pi"
	"fmt"
	"math/big"
	"strconv"
)

var (
	CORES = 4
	PI    = "314159265"
)

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

func consume(c chan string, join chan int) {
	for word := range c {
		n, _ := strconv.Atoi(word)
		if isPalindrome(word) && isPrime(int64(n)) {
			fmt.Println(word)
			join <- 1
		}
	}
}

func produce(c chan string, flag, prec int) {
	words := make(chan pi.Word)
	go pi.Chudnovsky(flag, prec, words)
	var current string = "141592653"
	var index int64 = 11
	for w := range words {
		for index < w.Digits-2 {
			if isPalindrome(current) {
				c <- current
			}
			current = current[1:] + string(w.Number[index])
			index++
		}
	}
}

func main() {
	words := make(chan string, 10) // channel of words within pi
	join := make(chan int)

	go produce(words, 100, 2000)
	for i := 0; i < CORES-1; i++ {
		go consume(words, join)
	}

	<-join
}
