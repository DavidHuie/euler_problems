package main

import "fmt"

func divisible(num int) bool {
	for j := 1; j <= 20; j++ {
		if num%j != 0 {
			return false
		}
	}
	return true
}

func problem5() int {
	var max int = 1
	for i := 1; i <= 20; i++ {
		max *= i
	}

	for i := 20; i < max; i += 20 {
		if divisible(i) {
			return i
		}
	}

	return 0
}

func main() {
	fmt.Println(problem5())
}
