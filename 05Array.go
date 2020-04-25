package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "a"
	names[1] = "b"
	names[2] = "c"

	fmt.Println(names[0], names[1], names[2])
	for i := 0; i < len(names); i++ { // len is Length
		fmt.Println(names[i])
	}
}
