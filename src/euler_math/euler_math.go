package euler_math

import "math"

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
