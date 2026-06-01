package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Jumlah angka:", len(numbers))
	fmt.Println("Index ke-1:", numbers[1])
	numbers[1] = 10
	fmt.Println("Isi array setelah diubah:", numbers)

	for index, value := range numbers {
		fmt.Println("Isi index ke :", index, " = ", value)
	}

	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // Membuat salinan array
	arr2[0] = 10
	fmt.Println("Isi array arr1:", arr1 == arr2) // arr1 tetap tidak berubah karena arr2 adalah salinan arr1
	fmt.Println("Isi array arr2:", arr2 != arr1) // arr2 berubah karena merupakan salinan arr1
}
