package main
import (
	"fmt"
	"net/http"
)

func quotesHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "server berjalan")
}

func main() {
	http.HandleFunc("/quotes", quotesHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Gagal menjalankan server", err)
	}
}