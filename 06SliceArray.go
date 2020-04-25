package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "a"
	names[1] = "b"
	names[2] = "c"

	fmt.Println(names[0], names[1], names[2]) // (output: a b c)
	var sliceArr = names[0:3]                 // For output On Array (output: [a b c])
	fmt.Println(sliceArr)
}
