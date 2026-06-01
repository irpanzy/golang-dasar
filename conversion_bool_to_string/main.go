package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Mengkonversi boolean ke string
	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println("Nilai boolean:", str)
	// Mengkonversi string ke boolean
	val, _ := strconv.ParseBool(str)
	fmt.Println("Nilai boolean setelah konversi balik:", val)
}
