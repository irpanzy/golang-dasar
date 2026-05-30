package main

import "fmt"

func main() {
	// Deklarasi variabel dengan tipe data eksplisit
	var name string = "Irpanzy"
	var age int = 22
	// Deklarasi variabel dengan tipe data implisit (type inference)
	city := "Yogyakarta"
	year := 2026
	fmt.Println("Nama:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Year:", year)
}
