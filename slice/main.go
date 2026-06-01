package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}

	s1 := arr[:] // Membuat slice yang mencakup seluruh array
	s2 := arr[2:5] // Membuat slice dari index 2 hingga 4 (5 tidak termasuk)
	fmt.Println("Isi slice:", s1)
	fmt.Println("Isi slice s2:", s2)

	s2[0] = 10
	fmt.Println("Isi slice s2 setelah diubah:", s2)
	fmt.Println("Isi array:", arr)
}