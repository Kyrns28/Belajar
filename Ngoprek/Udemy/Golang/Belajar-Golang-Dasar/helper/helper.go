package helper

// jika huruf kecil diawal tidak dapat akses dari luar package, namun jika huruf besar diawal dapat diakses dari luar package

// dapat diakses dari luar package
func SayHello(name string) string{
	return "Hello " + name
}

// Tidak dapat diakses dari luar package
func sayGoodbye(name string) string {
	return "Goodbye " + name
}