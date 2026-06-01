package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string = "85"
	number, err := strconv.Atoi(text) // Mengkonversi string ke int
	if err != nil {
		fmt.Println("Error konversi:", err)
	} else {
		fmt.Println("Nilai ujian:", number)
	}
}
