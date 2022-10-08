package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"
)

var (
	// INITIAL_START int64 = 0
	INITIAL_START    int64 = 1426403000
	LEN_WORD               = 21
	OPTIMAL_N_MILLIN       = 100
	REQUEST_PERIOD         = 110 * time.Millisecond
	TIMEOUT                = 35 * time.Second
)

type Response struct {
	Content string `json:"content"`
}

func getAndInc(c chan int64) int64 {
	out := <-c
	c <- out + 979
	return out
}

func bruteWay(n *big.Int) bool {
	mod2 := big.NewInt(0).Mod(n, big.NewInt(2))
	mod3 := big.NewInt(0).Mod(n, big.NewInt(3))
	if mod2.Int64() == 0 || mod3.Int64() == 0 {
		return false
	}
	root := big.NewInt(0).Sqrt(n).Int64()
	div := int64(5)
	for div <= root+1 {
		modF := big.NewInt(0).Mod(n, big.NewInt(div)).Int64()
		if modF == 0 {
			return false
		}
		modFPlus2 := big.NewInt(0).Mod(n, big.NewInt(div+2)).Int64()
		if modFPlus2 == 0 {
			return false
		}
		div += 6
	}
	return true
}

func isPrime(word string) bool { // function to check if number is prime
	n, _ := big.NewInt(0).SetString(word, 10)
	if n.ProbablyPrime(OPTIMAL_N_MILLIN) {
		// return bruteWay(n)
		return true
	}
	return false
}

func isPalindrome(s string) bool { // function to check if word is a palindrome
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func parseDigits(piDigits chan string, words chan string) {
	for digits := range piDigits {
		current := digits[:LEN_WORD]
		for i := LEN_WORD; i < 1000; i++ {
			current = current[1:] + string(digits[i])
			words <- current
		}
	}
}

func getPi(start int64) string {
	for {
		if start%(979*1000) == 0 {
			fmt.Println(start)
		}

		numberOfDigits := 1000
		radix := 10
		url := fmt.Sprintf("https://api.pi.delivery/v1/pi?start=%v&numberOfDigits=%v&radix=%v", start, numberOfDigits, radix)

		raw, err := http.Get(url)

		if err != nil {
			log.Fatalf("Error while requesting digits: %v", err)
		}

		defer raw.Body.Close()

		switch raw.StatusCode {
		case 200:
			jsonResp, _ := ioutil.ReadAll(raw.Body)

			var resp Response
			json.Unmarshal(jsonResp, &resp)

			return resp.Content
		default:
			log.Fatalf("Unexpected status code: %v", raw.Status)
		}
	}
}

func produce(words chan string) {
	piDigits := make(chan string, 10)
	startChan := make(chan int64, 1)
	startChan <- INITIAL_START
	go parseDigits(piDigits, words)
	tick := time.NewTicker(REQUEST_PERIOD)
	for range tick.C {
		go func() {
			start := getAndInc(startChan)
			piDigits <- getPi(start)
		}()
	}
}

func consume(words chan string, join chan int) {
	for word := range words {
		if isPalindrome(word) && isPrime(word) {
			fmt.Println(word)
			join <- 1
		}
	}
}

func main() {
	words := make(chan string, 10)
	join := make(chan int)

	t := time.Now()

	go produce(words)
	go consume(words, join)

	<-join

	fmt.Printf("Done in time %v\n", time.Since(t))
}
