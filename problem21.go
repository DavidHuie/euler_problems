package main

import "fmt"
import "euler_math"

func problem21() int {
	number_checked := make(map[int]bool)
	sum := 0

	for i := 1; i < 10000; i++ {
		if number_checked[i] {
			continue
		}

		sum_of_divisors := euler_math.SumOfDivisors(i)

		if sum_of_divisors == i {
			continue
		}

		if euler_math.SumOfDivisors(sum_of_divisors) == i {
			number_checked[sum_of_divisors] = true
			sum += i + sum_of_divisors
		}
	}

	return sum
}

func main() {
	fmt.Println(problem21())
}
