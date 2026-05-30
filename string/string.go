package main

import "fmt"

func main() {
<<<<<<< HEAD
	// Deklarasi string dengan menggunakan double quotes
	nama := "Irpanzy"
	pesan := "Selamat belajar Go!"

	// Deklarasi string dengan menggunakan backticks (raw string literal)
=======

	nama := "Irpanzy"
	pesan := "Selamat belajar Go!"

>>>>>>> 359184e8d116f71c9db801509f2e894b5ff5793b
	paragraf := `Halo, nama saya ` + nama + `.
` + pesan + `
Semoga hari Anda menyenangkan!`

<<<<<<< HEAD
	// Menampilkan string
=======
>>>>>>> 359184e8d116f71c9db801509f2e894b5ff5793b
	fmt.Println(paragraf)
}
