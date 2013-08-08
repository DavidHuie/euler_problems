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
	if num%2 == 0 {
		return false
	}

	var max int = int(math.Sqrt(float64(num))) + 1

	for i := 2; i < max; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// import "sync"
// var wait_group sync.WaitGroup
// wait_group.Add(1)
