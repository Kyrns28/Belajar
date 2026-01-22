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

	// Example assignment operator
	f := b
	f += 3

	// Example Comparison Operator
	x := 5
	y := 3

	fmt.Println(x>y)
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
	fmt.Println("F = ", f)

}

// List Assignment Operator
// =	x = 5	x = 5	
// +=	x += 3	x = x + 3	
// -=	x -= 3	x = x - 3	
// *=	x *= 3	x = x * 3	
// /=	x /= 3	x = x / 3	
// %=	x %= 3	x = x % 3	
// &=	x &= 3	x = x & 3	
// |=	x |= 3	x = x | 3	
// ^=	x ^= 3	x = x ^ 3	
// >>=	x >>= 3	x = x >> 3	
// <<=	x <<= 3	x = x << 3

// List Comparison operator

// Operator		Name						Example	
// ==			Equal to					x == y	
// !=			Not equal					x != y	
// >			Greater than				x > y	
// <			Less than					x < y	
// >=			Greater than or equal to	x >= y	
// <=			Less than or equal to		x <= y

// List Logical Operator
// Operator	Name			Description													Example	
// && 			Logical and		Returns true if both statements are true					x < 5 &&  x < 10	
// || 			Logical or		Returns true if one of the statements is true				x < 5 || x < 4	
// !			Logical not		Reverse the result, returns false if the result is true		!(x < 5 && x < 10)