package main
import "fmt"

func main() {
	// const firstName string = "Rizki"
	// const middleName = "Nur"

	// fmt.Println(firstName)
	// fmt.Println(middleName)

	// error karena const tidak dapat diubah lagi
	// firstName = "Tes"

	// Multiple Const
	const (
		firstName string = "Rizki"
		middleName = "Nur"
		LastName = "Sepriana"
	)

	fmt.Println (firstName)
	fmt.Println(middleName)
	fmt.Println(LastName)
}