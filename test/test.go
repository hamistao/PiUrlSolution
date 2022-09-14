package main

import (
	"challenge/pi"
)

func binarySplitFactorial(a, b int) int {
	return 0
}

func incIndex(ch chan int) int {
	i := <-ch
	ch <- i + 1
	return i
}

func main() {
	// words := make(chan pi.Word)
	// go pi.Chudnovsky(50, 1000, words)
	// var current string = "141592653"
	// var index int64 = 11
	// for w := range words {
	// 	for index < w.Digits {
	// 		fmt.Println(current)
	// 		current = current[1:] + string(w.Number[index])
	// 		index++
	// 	}
	// }

	pi.CalcPi(1489)
}
