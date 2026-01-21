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
	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(sum3)
}
