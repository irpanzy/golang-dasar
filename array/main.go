package main

import "fmt"

func main() {
	var angka [3]int = [3]int{1, 2, 3} // Mendeklarasikan array dengan tipe data int dan panjang 3
	fmt.Println("Isi array angka:", angka)
	fmt.Println(angka[1])
	angka[1] = 20
	fmt.Println("Isi array angka setelah diubah:", angka)
}
