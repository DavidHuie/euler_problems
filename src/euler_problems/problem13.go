package main

import "fmt"
import "io/ioutil"
import "strings"
import "math/big"

func main() {
	file, _ := ioutil.ReadFile("./problem13.txt")
	num_str := string(file[:len(file)])
	trimmed_num_str := strings.TrimSpace(num_str)
	lines := strings.Split(trimmed_num_str, "\n")

	sum := big.NewInt(0)

	for _, line := range lines {
		i := big.NewInt(0)
		i.SetString(line, 10)
		sum.Add(sum, i)
	}

	sum_string := sum.String()
	sum_string = sum_string[:10]

	fmt.Println(sum_string)
}
