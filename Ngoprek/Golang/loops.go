package main
import ("fmt")

func loopsGo(){
	// example 1
	for i := 0; i < 5 ; i++ {
		fmt.Println(i)
	}

	// example 2
	for i := 0; i<= 100; i+=10{
		fmt.Println(i)
	}

	// example 3
	for i := 0; i<5; i++ {
		if i ==3 {
			continue
		}
		fmt.Println(i)
	} // output 0,1,2,4

	// example 4
		for i := 0; i<5; i++ {
		if i ==3 {
			break
		}
		fmt.Println(i)
	} // output 0,1,2
}