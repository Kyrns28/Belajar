package main

import "fmt"

// "type Todo struct" mendefinisikan tipe data baru bernama Todo
// struct adalah kumpulan field yang membentuk satu kesatuan data
type Todo struct {
	// field pertama: nama task, bertipe string
	Title string
	// field kedua: status selesai atau belum, bertipe bool (true/false)
	Done bool
}

func main() {
	// slice kosong yang isinya bertipe Todo (bukan string biasa)
	// []Todo artinya "slice yang isinya struct Todo"
	var todos []Todo

	// loop selamanya sampai user memilih keluar
	for {
		fmt.Println("\nTo do list CLI")
		fmt.Println("1. Tambah task")
		fmt.Println("2. Lihat semua task")
		fmt.Println("3. Tandai task selesai")
		fmt.Println("4. Hapus task")
		fmt.Println("5. Keluar")

		// deklarasi di dalam loop agar setiap iterasi selalu mulai dari 0
		var pilihan int
		fmt.Print("Pilih Menu (1-5): ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1: // ── TAMBAH TASK ──────────────────────────────────────
			var title string
			fmt.Print("Masukan nama Task : ")
			fmt.Scan(&title)

			// membuat instance baru dari struct Todo — disebut "struct literal"
			// setiap field diisi sesuai namanya
			task := Todo{
				Title: title, // diisi dari input user
				Done:  false, // task baru selalu belum selesai
			}
			// menambahkan struct task ke slice todos
			todos = append(todos, task)
			fmt.Println("Task berhasil ditambahkan")

		case 2: // ── LIHAT SEMUA TASK ─────────────────────────────────
			// len() mengembalikan jumlah elemen dalam slice
			// jika 0 berarti slice masih kosong
			if len(todos) == 0 {
				fmt.Println("Belum ada task")
				// break keluar dari switch, kembali ke awal for loop
				break
			}
			// iterasi slice of struct
			// "todo" adalah satu elemen Todo, bisa akses todo.Title dan todo.Done
			for i, todo := range todos {
				status := "❌"
				// karena Done bertipe bool, tidak perlu "== true"
				// cukup "if todo.Done" saja
				if todo.Done {
					status = "✅"
				}
				fmt.Printf("%d. [%s] %s\n", i+1, status, todo.Title)
			}

		case 3: // ── TANDAI TASK SELESAI ──────────────────────────────
			if len(todos) == 0 {
				fmt.Println("Belum ada task")
				break
			}
			// tampilkan task dulu agar user tahu nomor yang ingin ditandai
			for i, todo := range todos {
				status := "❌"
				if todo.Done {
					status = "✅"
				}
				fmt.Printf("%d. [%s] %s\n", i+1, status, todo.Title)
			}

			var nomor int
			fmt.Print("Masukan nomor task yang selesai: ")
			fmt.Scan(&nomor)

			// validasi dua sisi:
			// nomor < 1 → terlalu kecil (tidak ada task nomor 0 atau negatif)
			// nomor > len(todos) → terlalu besar (melebihi jumlah task)
			// "||" artinya "atau" — salah satu true = kondisi true
			if nomor < 1 || nomor > len(todos) {
				fmt.Println("Nomor task tidak valid")
				break
			}
			// akses elemen slice di index nomor-1
			// (-1 karena index mulai dari 0, tapi user input mulai dari 1)
			// lalu ubah field Done menjadi true
			todos[nomor-1].Done = true
			fmt.Println("Task sudah selesai")

		case 4: // ── HAPUS TASK ───────────────────────────────────────
			if len(todos) == 0 {
				fmt.Println("Belum ada task")
				break
			}
			// tampilkan task dulu agar user tahu nomor yang ingin dihapus
			for i, todo := range todos {
				status := "❌"
				if todo.Done {
					status = "✅"
				}
				fmt.Printf("%d. [%s] %s\n", i+1, status, todo.Title)
			}

			var nomor int
			fmt.Print("Masukan nomor task yang ingin dihapus: ")
			fmt.Scan(&nomor)

			if nomor < 1 || nomor > len(todos) {
				fmt.Println("Nomor task tidak valid")
				break
			}
			// cara menghapus elemen dari slice:
			// todos[:nomor-1] → semua elemen SEBELUM yang dihapus
			// todos[nomor:]   → semua elemen SETELAH yang dihapus
			// "..." → unpack slice menjadi elemen individual untuk append
			// hasilnya: slice baru tanpa elemen yang dihapus
			todos = append(todos[:nomor-1], todos[nomor:]...)
			fmt.Println("Task berhasil dihapus!")

		case 5: // ── KELUAR ───────────────────────────────────────────
			fmt.Println("Sampai jumpa!")
			// return keluar dari func main = program selesai sepenuhnya
			// lebih tepat dari break untuk keluar dari loop di dalam main
			return

		default:
			fmt.Println("Angka yang dimasukan tidak valid")
		}
	}
}