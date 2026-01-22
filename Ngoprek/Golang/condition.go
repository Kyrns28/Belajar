package main
import ("fmt")

func condition(){

	// example if
	// x := 20
	// y := 15

	// if x > y {
	// 	fmt.Println("x is greater than y")
	// }

	// Example if else
	// temperature := 16

	// if (temperature > 15){
	// 	fmt.Println("it is warm out here")
	// } else {
	// 	fmt.Println("it is cool out here")
	// }

	// Example else if
		// x := 30
		// if x >= 10 {
		// 	fmt.Println("x is larger than or equal 10")
		// } else if x >20 {
		// 	fmt.Println("x is larger than 20")
		// } else {
		// 	fmt.Println("x is less than 20")
		// }


	// Example Nested IF
	num := 20
	
	if num >= 10 {
		fmt.Println("Num is more than 10")
		if num > 15{
			fmt.Println("Num is Also more than 15")
		}
	} else {
		fmt.Println("Num is less than 10")
	}
}