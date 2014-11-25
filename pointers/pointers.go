package main

import "fmt"

/*
We’ll show how pointers work in contrast to values with 2 functions:
  zeroval and zeroptr.
*/

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("Pointer: ", &i)
}
