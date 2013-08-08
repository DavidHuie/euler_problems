package euler_math

import "math"
import "container/list"

func FibList(num uint) []uint {
	var fib_list []uint = make([]uint, num)

	fib_list[0] = 0
	fib_list[1] = 1

	for i := 2; i < len(fib_list); i++ {
		fib_list[i] = fib_list[i-1] + fib_list[i-2]
	}

	return fib_list
}

func IsPrime(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	ch := make(chan bool, 1000000)

	divisible := func(num int, d int, ch chan bool) {
		ch <- (num%d == 0)
	}

	var max int = int(math.Sqrt(float64(num))) + 1

	for i := 3; i < max; i += 2 {
		go divisible(num, i, ch)
	}

	for i := 3; i < max; i += 2 {
		if <-ch {
			return false
		}
	}

	return true
}

// PrimesLessThan helpers

func n_to_i(n int) int {
	return n - 2
}

func i_to_n(i int) int {
	return i + 2
}

// Returns a slice containing the primes less than num. This
// uses the Sieve of Eratosthenes algorithm.
func PrimesLessThan(num int) []int {
	primes := list.New()
	slice_size := n_to_i(num) + 1
	possible_factors := make([]bool, slice_size)

	for i, _ := range possible_factors {
		possible_factors[i] = true
	}

	for i, is_prime := range possible_factors {
		if !is_prime {
			continue
		}
		for j := i; j < slice_size; j += i_to_n(i) {
			possible_factors[j] = false
		}
		primes.PushBack(i_to_n(i))
	}

	primes_slice := make([]int, primes.Len())
	index := 0
	for prime := primes.Front(); prime != nil; prime = prime.Next() {
		primes_slice[index] = prime.Value.(int)
		index++
	}

	return primes_slice
}

func SumOfSquares(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i * i
	}

	return sum
}

func SquareOfSum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}

	return sum * sum
}

// import "sync"
// var wait_group sync.WaitGroup
// wait_group.Add(1)
