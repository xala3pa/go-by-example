package main

import (
	"fmt"
)

func plus(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	fmt.Println("sum: ", plus(2, 5))
}
