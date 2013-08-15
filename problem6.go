package main

import "fmt"
import "./euler_math"

func problem6() int {
	return euler_math.SquareOfSum(100) - euler_math.SumOfSquares(100)
}

func main() {
	fmt.Println(problem6())
}
