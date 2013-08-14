package main

import "fmt"
import "euler_math"

func problem23() int {
	max_int := 28123

	// Locate all abundant numbers under the limit
	abundant_map := make(map[int]bool)
	for i := 1; i <= max_int; i++ {
		if euler_math.SumOfDivisors(i) > i {
			abundant_map[i] = true
		}
	}

	// An integer is special if it is made up of
	// two abundant numbers.
	special_map := make(map[int]bool)
	for i := 1; i <= max_int; i++ {
		i_is_abundant := abundant_map[i]
		if i_is_abundant {
			for j := i; j <= max_int; j++ {
				j_is_abundant := abundant_map[j]
				if j_is_abundant {
					special_map[i + j] = true
				}
			}
		}
	}

	sum := 0
	for i := 1; i <= max_int; i++ {
		if !special_map[i] {
			sum += i
		}
	}

	return sum
}

func main() {
	fmt.Println(problem23())
}
