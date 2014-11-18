package main

import (
	"fmt"
)

func main() {

	//Also knows in other languages as dicts or hashes

	m := make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2
	fmt.Println("map: ", m)

	v1 := m["key1"]
	fmt.Println("get v1: ", v1)

	fmt.Println("len: ", len(m))

	delete(m, "key2")
	fmt.Println("map: ", m)

	_, prs := m["key2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("literal map: ", n)
}
