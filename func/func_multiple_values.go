package main

import (
	"fmt"
)

func multi_values() (int, int) {
	return 7, 9
}

func main() {
	a, b := multi_values()
	fmt.Println(a)
	fmt.Println(b)

	_, c := multi_values()
	fmt.Println(c)
}
