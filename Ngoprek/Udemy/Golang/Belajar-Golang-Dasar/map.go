package main
import "fmt"

func main(){
	person := map[string]string{
		"name" : "Rizkinur",
		"address" : "Depok",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)


	book := make(map[string]string)
	book["title"] = "Buku golang"
	book["author"] = "Rizkinur"
	book["upps"] = "Salah"

	fmt.Println(book)
	
	delete(book, "upps")
	fmt.Println(book)
	
}