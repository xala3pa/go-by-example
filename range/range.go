package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		num += num
	}
	fmt.Println("sum: ", sum)

	for index, num := range nums {
		if num == 3 {
			fmt.Println("index: ", index)
		}
	}

	kvs := map[string]string{"Alvaro": "37", "Sandra": "36", "Hector": "4"}
	for k, v := range kvs {
		fmt.Printf("%s is %s years old\n", k, v)
	}

	/*range on strings iterates over Unicode code points.
	The first value is the starting byte index of the rune and the second the rune itself.*/
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
