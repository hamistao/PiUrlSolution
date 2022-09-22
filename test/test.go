package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

type Response struct {
	Content string `json:"content"`
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func checkDivisors(candidate *big.Int, join chan int, out chan bool, i, j int64) {
	fmt.Print("thread: ")
	fmt.Println(i)
	var count int64 = 0
	for i <= j {
		if count%50000000 == 0 {
			fmt.Print("div:")
			fmt.Println(i)
		}
		modF := big.NewInt(0).Mod(candidate, big.NewInt(i)).Int64()
		if modF == 0 {
			fmt.Println("foi porra")
			out <- false && <-out
			join <- 1
			return
		}
		modFPlus2 := big.NewInt(0).Mod(candidate, big.NewInt(i+2)).Int64()
		if modFPlus2 == 0 {
			fmt.Println("foi porra")
			out <- false && <-out
			join <- 1
			return
		}
		i += 6
		count++
	}
}

func bruteWay(n *big.Int, n_threads int) bool {
	mod2 := big.NewInt(0).Mod(n, big.NewInt(2))
	fmt.Println("foi mod2")
	mod3 := big.NewInt(0).Mod(n, big.NewInt(3))
	fmt.Println("foi mod3")
	if mod2.Int64() == 0 || mod3.Int64() == 0 {
		return false
	}
	root := big.NewInt(0).Sqrt(n).Int64()
	div := int64(5)
	step := int64((root-div)/int64(n_threads)) + 2
	fmt.Print("step: ")
	fmt.Println(step)
	join := make(chan int)
	out := make(chan bool, 1)
	fmt.Println("a")
	out <- true

	fmt.Println("a")
	for i := 0; i < n_threads; i++ {
		go checkDivisors(n, join, out, div, min(div+step, root))
		div += step
	}

	for i := 0; i < n_threads; i++ {
		<-join
	}

	return <-out
}

func isPrime(word string) bool { // function to check if number is prime
	n, _ := big.NewInt(0).SetString(word, 10)
	if n.ProbablyPrime(0) {
		fmt.Println("probably")
		return bruteWay(n, 10)
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

func getPi(start int) string {
	numberOfDigits := 1000
	radix := 10
	url := fmt.Sprintf("https://api.pi.delivery/v1/pi?start=%v&numberOfDigits=%v&radix=%v", start, numberOfDigits, radix)

	raw, _ := http.Get(url)

	jsonResp, _ := ioutil.ReadAll(raw.Body)

	var resp Response
	json.Unmarshal(jsonResp, &resp)

	return resp.Content
}

func main() {
	// word := "100000000000000016471"
	// word := "923794682393286497323"
	// n, _ := big.NewInt(0).SetString(word, 10)
	// fmt.Println(bruteWay(n))

	// fmt.Println(isPrime(word))

	fmt.Println(getPi(2937979))
}
