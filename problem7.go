package main

import "fmt"
import "./euler_math"

func problem7() int {
	prime_count := 0
	for i := 1; ; i++ {
		if euler_math.IsPrime(i) {
			prime_count++
		}
		if prime_count == 10001 {
			return i
		}
	}
}

func main() {
	fmt.Println(problem7())
}
