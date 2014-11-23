package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(2, 3, 4)

	/*If you already have multiple args in a slice,
	apply them to a variadic function using func(slice...) like this.*/
	nums := []int{1, 3, 5}
	sum(nums...)
}
