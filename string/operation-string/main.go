package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, World!"

	// Mengubah semua huruf menjadi huruf kecil
	fmt.Println("Lowercase:", strings.ToLower(text))

	// Mengubah semua huruf menjadi huruf besar
	fmt.Println("Uppercase:", strings.ToUpper(text))

	// Mengecek apakah string dimulai dengan "Hello"
	fmt.Println("Starts with 'Hello':", strings.HasPrefix(text, "Hello"))

	// Mengecek apakah string mengandung kata "World"
	fmt.Println("Contains 'World':", strings.Contains(text, "World"))

	// Memisahkan string berdasarkan delimiter
	parts := strings.Split("apele, jeruk, pisang", ", ")
	fmt.Println("Split:", parts)
}
