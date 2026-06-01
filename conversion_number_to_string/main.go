package main

import (
	"fmt"
	"strconv"
)

func main() {
	var score int = 85
	var text string = strconv.Itoa(score) // Mengkonversi int ke string

	fmt.Println("Nilai ujian:", text)
}
