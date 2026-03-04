package main
import (
	"fmt"
	"os"
	"strings"
	"math/rand"
	"time"
)

func main() {
	data, err := os.ReadFile("quotes.txt")

	if err != nil {
		fmt.Println("Gagal menjalankan file", err)
		return
	}
	content := string(data)
	quotes := strings.Split(content, "\n")


	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(quotes))

	quote := quotes[index]

	fmt.Println("Quote hari ini : ", quote)
}