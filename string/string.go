package main

import "fmt"

func main() {
	// Deklarasi string dengan menggunakan double quotes
	nama := "Irpanzy"
	pesan := "Selamat belajar Go!"

	// Deklarasi string dengan menggunakan backticks (raw string literal)
	paragraf := `Halo, nama saya ` + nama + `.
` + pesan + `
Semoga hari Anda menyenangkan!`

	// Menampilkan string
	fmt.Println(paragraf)
}
