package main

import "fmt"
import "math"
import "./euler_math"

func problem3() int {
	var largest_prime_factor int
	var number int = 600851475143
	var max_factor int = int(math.Sqrt(float64(number)))

	for i := 2; i < max_factor; i++ {
		if (number % i == 0) && (euler_math.IsPrime(i)) {
			largest_prime_factor = i
		}
	}

	return largest_prime_factor
}

func main() {
	fmt.Println(problem3())
}
