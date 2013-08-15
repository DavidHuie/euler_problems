package main

import "fmt"
import "./euler_math"

func problem2() uint {
	var sum uint

	fib_list := euler_math.FibList(4000000)
	for _, fib := range fib_list {
		if fib > 4000000 {
			break
		}
		if fib%2 == 0 {
			sum += fib
		}
	}

	return sum
}

func main() {
	fmt.Println(problem2())
}
