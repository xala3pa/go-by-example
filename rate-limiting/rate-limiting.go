package main

import "time"
import "fmt"

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 500)

	/*
	 * By blocking on a receive from the limiter channel
	 * before serving each request, we limit ourselves
	 * to 1 request every 200 milliseconds.
	 */
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

}
