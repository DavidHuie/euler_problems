package main

import "fmt"
import "euler_math"

func problem10() int {
	num := 2000000
	sum := 0
	for _, val := range euler_math.PrimesLessThan(num) {
		sum += val
	}

	return sum
}

func main() {
	fmt.Println(problem10())
}
