package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"sync"
	"time"
)

var (
	N_WORKERS = 100
	LEN_WORD  = 17
	N_MILLIN  = 100
	FILE_PATH = "/home/pedro/pi0.txt" //https://storage.googleapis.com/pi100t/Pi%20-%20Dec%20-%20Chudnovsky/Pi%20-%20Dec%20-%20Chudnovsky%20-%201.ycd
)

func isPrime(word string) bool { // function to check if number is prime
	n, _ := big.NewInt(0).SetString(word, 10)
	return n.ProbablyPrime(N_MILLIN)
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

func sweepDigits(digits string) (string, string, bool) {
	initialValue := string(digits[:LEN_WORD])
	current := initialValue
	for i := LEN_WORD; i < len(digits); i++ {
		current = current[1:] + string(digits[i])
		if isPalindrome(current) && isPrime(current) {
			return "", current, true
		}
	}

	return initialValue, current, false
}

func produce(join chan struct{}) {
	file, err := os.Open(FILE_PATH)

	if err != nil {
		fmt.Println("cannot open the file", err)
		return
	}

	defer file.Close() //close after checking err

	r := bufio.NewReader(file)

	chunkPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 500*1024)
		return &lines
	}}

	stringPool := sync.Pool{New: func() interface{} {
		word := ""
		return &word
	}}

	waitChan := make(chan struct{}, N_WORKERS)
	mutex := make(chan struct{}, 1)
	id := 0
	var barrier sync.WaitGroup

	leftOvers := make(map[int]string)

	for {
		buf := *chunkPool.Get().(*[]byte)
		n, err := r.Read(buf)
		buf = buf[:n]
		if n == 0 {
			if err != nil {
				fmt.Println(err)
				break
			}
			if err == io.EOF {
				break
			}
			return
		}

		waitChan <- struct{}{}
		barrier.Add(1)
		go func(chunk []byte, id int) {
			digits := *stringPool.Get().(*string)
			digits = string(chunk)
			chunkPool.Put(&chunk)
			init, final, ok := sweepDigits(digits)
			stringPool.Put(&digits)
			if ok {
				fmt.Println(final)
				join <- struct{}{}
				return
			}

			mutex <- struct{}{}
			leftOvers[id] += init
			leftOvers[id+1] += final
			<-mutex

			barrier.Done()
			<-waitChan
		}(buf, id)
		id++
	}
	barrier.Wait()
	processLeftOvers(leftOvers, id, join)
	join <- struct{}{}
}

func processLeftOvers(leftOvers map[int]string, max int, join chan struct{}) {
	for i, digits := range leftOvers {
		if i != 0 && i != max {
			_, final, ok := sweepDigits(digits)
			if ok {
				fmt.Println(final)
				join <- struct{}{}
				return
			}
		}
	}
}

func main() {
	join := make(chan struct{})
	t := time.Now()

	go produce(join)

	<-join
	fmt.Printf("Done in time %v\n", time.Since(t))
}
