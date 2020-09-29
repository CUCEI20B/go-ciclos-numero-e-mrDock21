package main

import (
	"fmt"
	"strconv"
)

func factorial(number int) float32 {
	var res float32 = 0

	if number > 1 {
		res = float32(number) * float32(number - 1)
		for i := number - 2; i > 0; i-- {
			res = res * float32(i)
		}
	} else {
		res = 1
	}

	return res
}

func main() {
	var input string
	var res float32 = 0

	// Read from console: number of iterations
	fmt.Scan(&input)

	// Scan values from input string
	N, _ := strconv.Atoi(input)

	for i := 0; i <= N; i++ {
		fact := factorial(i)
		res += float32(1) / fact
	}

	fmt.Println(res)
}