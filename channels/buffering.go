package main

import "fmt"

func main() {
	meessages := make(chan string, 2)

	meessages <- "buffered"
	meessages <- "channel"

	fmt.Println(<-meessages)
	fmt.Println(<-meessages)
}
