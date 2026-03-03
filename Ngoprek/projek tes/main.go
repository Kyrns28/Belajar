package main
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("quotes.txt")

	if err != nil {
		fmt.Println("Gagal menjalankan file", err)
		return
	}

	// fmt.Println(string(data))

	content := string(data)
	quotes := strings.Split(content, "\n")

	for i, quote := range quotes{
		fmt.Printf("%d. %s\n", i+1, quote)
	}
}