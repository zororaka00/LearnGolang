package main

import "fmt"

func main() {
	var a int = 0
	fmt.Println(&a)

	var b [2]int8
	b[0] = 1
	b[1] = 2
	fmt.Println(&b[0], &b[1])
}
