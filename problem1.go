package main

import "fmt"

func problem1() int {
	var sum int
	for i := 1; i < 1000; i++ {
		if (i % 3 == 0) || (i % 5 == 0) {
			sum += i
		}
	}

	return sum
}

func main() {
	fmt.Println(problem1())
}
