package main

import "fmt"

func main() {
	var varstring string = "Check"
	varstring = "Check Change"
	fmt.Println(varstring)

	// uint8 (0 - 255)
	// uint16 (0 - 65535)
	// uint32 (0 - 4294967295)
	// uint64 (0 - 18446744073709551615)
	var varuint uint = 18446744073709551615 // like uint64
	fmt.Println(varuint)

	var varbyte byte = 255 // like uint8 (0 - 255)
	fmt.Println(varbyte)

	// int8 (-128 - 127)
	// int16 (-32768 - 32767)
	// int32 (-2147483648 - 2147483647)
	// int64 (-9223372036854775808 - 9223372036854775807)
	var varint int = -9223372036854775808 // like int64
	fmt.Println(varint)

	var varrune rune = -2147483648 // like int32 (-2147483648 - 2147483647)
	fmt.Println(varrune)

	// float32 (-3.4e+38 - 3.4e+38)
	// float64 (-1.7e+308 - 1.7e+308)
	var vardecimal float64 = 1.7e+308
	fmt.Println(vardecimal)

	var varbool bool = true
	fmt.Println(varbool)
}
