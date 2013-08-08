package euler_math

func FibList(num uint) []uint {
	var fib_list []uint = make([]uint, num)

	fib_list[0] = 0
	fib_list[1] = 1

	for i := 2; i < len(fib_list); i++ {
		fib_list[i] = fib_list[i-1] + fib_list[i-2]
	}

	return fib_list
}
