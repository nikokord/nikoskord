package main

import (
	"fmt"
)

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 {
		return 1
	}

	arg := 1

	for i := 2; i <= nb; i++ {
		arg *= i
	}

	return arg
}

func main() {
	fmt.Println(IterativeFactorial(-2))
}
