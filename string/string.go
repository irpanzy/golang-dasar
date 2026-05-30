package main

import "fmt"

func main() {

	nama := "Irpanzy"
	pesan := "Selamat belajar Go!"

	paragraf := `Halo, nama saya ` + nama + `.
` + pesan + `
Semoga hari Anda menyenangkan!`

	fmt.Println(paragraf)
}
