package main

import "fmt"
import "./euler_math"

func problem14() int {
	length := 0
	num := 0

	for i := 1; i <= 1000000; i++ {
		new_length := euler_math.CollatzSequenceLength(i)
		if new_length > length {
			length = new_length
			num = i
		}
	}

	return num
}

func main() {
	fmt.Println(problem14())
}
