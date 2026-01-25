package main
import "fmt"

func main(){
	counter := 1

	for counter <= 10 {
		fmt.Println("Ini perulangan yang ke ", counter)
		counter++
	}

	// for dengan statement
	for counter := 1; counter <=10; counter++ {
		fmt.Println("Ini perulangan yang ke ", counter)
	}
}