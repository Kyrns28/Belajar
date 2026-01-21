package main

import (
	"fmt"
)

func operators() {
	var a = 15 + 25
	fmt.Println("a = ", a)

	var (
		sum1 = 100 + 150
		sum2 = sum1 + 200
		sum3 = sum2 + 500
	)
	fmt.Println("Sum1 = ", sum1)
	fmt.Println("Sum2 = ", sum2)
	fmt.Println("Sum3 = ", sum3)
}
