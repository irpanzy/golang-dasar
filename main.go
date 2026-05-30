package main

import "fmt"

func main() {
	// Deklarasi variabel dengan tipe data eksplisit
	var name string = "Irpanzy"
	var age int = 22

	// Deklarasi variabel dengan tipe data implisit (type inference)
	city := "Yogyakarta"
	year := 2026

	// Deklarasi konstanta
	const pi = 3.14
	const appName = "Learning Go"

	fmt.Println("Nama:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Year:", year)
	fmt.Println("Pi:", pi)
	fmt.Println("App Name:", appName)

	pi = 3.14159 // Error: cannot assign to pi (constant)
}
