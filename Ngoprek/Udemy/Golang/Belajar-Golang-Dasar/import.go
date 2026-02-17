package main
import (
	"Belajar-Golang-Dasar/helper"
	"fmt"
)

func main(){
	result := helper.SayHello("Kinur")
	fmt.Println(result)

	// kode normal
	fmt.Println(helper.SayHello("Rizkinur"))

	// tidak dapat memanggil func nya karena huruf kecil diawal dan diluar package
	fmt.Println(helper.sayGoodbye("Rizkinur"))
}