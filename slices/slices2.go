package main

import (
	"fmt"
)

func main() {
	//An array type definition specifies a length and an element type.
	var a [5]int
	a[0] = 1
	i := a[0]
	fmt.Println("i value: ", i)
	//Arrays do not need to be initialized explicitly;
	fmt.Println("a[2] value: ", a[2] == 0)
	//Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element.

	//An array literal
	literalArr := [3]int{1, 2, 3}
	fmt.Println("literal array: ", literalArr)
	//compiler count the array elements
	lang := [...]string{"golang", "Dart"}
	fmt.Println("lang: ", lang)

	/*The slice type is an abstraction built on top of Go's array type.
	A slice literal is declared just like an array literal, except you leave out the element count*/
	letters := []string{"a", "b", "c", "d"}
	fmt.Println("letters: ", letters)
	//A slice can be created with the built-in function called make: func make([]T, len, cap) []T
	var s []byte
	s = make([]byte, 5, 5)
	fmt.Println("slice: ", s)
	//When the capacity argument is omitted, it defaults to the specified length
	s2 := make([]byte, 5)
	fmt.Println("slice2: ", s2)
	//The length and capacity of a slice can be inspected using the built-in len and cap functions.
	fmt.Println("Length: ", len(s) == 5)
	fmt.Println("Capacity: ", cap(s) == 5)
	fmt.Println(cap(s) == cap(s2))
	//The zero value of a slice is nil. The len and cap functions will both return 0 for a nil slice

	//A slice can also be formed by "slicing" an existing slice or array
	b := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Println(b)
	fmt.Println("b[:2] == []byte{'g', 'o'} ")
	fmt.Println("b[2:] == []byte{'l', 'a', 'n', 'g'}")
	fmt.Println("b[:] == b")
	/*A slice is a descriptor of an array segment. It consists of a pointer to the array,
	the length of the segment, and its capacity (the maximum length of the segment).
	Slicing does not copy the slice's data. It creates a new slice value that points to the original array.
	A slice cannot be grown beyond its capacity.
	To append one slice to another, use ... to expand the second argument to a list of arguments.*/
	x := []string{"John", "Paul"}
	y := []string{"George", "Ringo", "Pete"}
	x = append(x, y...) // equivalent to "append(a, b[0], b[1], b[2])"
	fmt.Println("Append...:", x)

}
