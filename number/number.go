package main

import (
	"fmt"
)

func main() {
	// Signed integers
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = -100 // Ukuran tergantung pada platform (32-bit atau 64-bit)

	// Unsigned integers
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint = 100

	// Print the values
	fmt.Println("Signed Integers:")
	fmt.Printf("int8	: %d\n", i8)
	fmt.Printf("int16	: %d\n", i16)
	fmt.Printf("int32	: %d\n", i32)
	fmt.Printf("int64	: %d\n", i64)
	fmt.Printf("int	: %d\n", i)

	// Print the values
	fmt.Println("Unsigned Integers:")
	fmt.Printf("uint8	: %d\n", u8)
	fmt.Printf("uint16	: %d\n", u16)
	fmt.Printf("uint32	: %d\n", u32)
	fmt.Printf("uint64	: %d\n", u64)
	fmt.Printf("uint	: %d\n", u)
}
