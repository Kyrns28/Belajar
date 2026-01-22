package main

import (
	"fmt"
)

func operators() {
	var a = 15 + 25
	var b = 15 - 10
	c := 15 * 2
	d := 16 / 2
	e := 17 % 2 

	var (
		sum1 = 100 + 150
		sum2 = sum1 + 200
		sum3 = sum2 + 500
	)
	fmt.Println("Sum1 = ", sum1)
	fmt.Println("Sum2 = ", sum2)
	fmt.Println("Sum3 = ", sum3)
	fmt.Println("A = ", a)
	fmt.Println("B = ", b)
	fmt.Println("C = ", c)
	fmt.Println("D = ", d)
	fmt.Println("D = ", e)

}
