package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Alvaro", 37})

	fmt.Println(person{name: "Sandra", age: 35})

	fmt.Println(person{name: "Luis"})

	fmt.Println(&person{name: "Puri", age: 59})

	s := person{name: "Luis", age: 23}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.name)

	sp.age = 90
	fmt.Println(sp.age)
}
