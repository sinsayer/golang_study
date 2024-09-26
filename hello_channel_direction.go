package main

import (
	"fmt"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	fmt.Println("start")
	pings := make(chan string, 2)
	fmt.Println("before ping")
	pongs := make(chan string, 2)
	fmt.Println("before pong")

	fmt.Println("go pong")
	go pong(pings, pongs)
	fmt.Println("after pong")

	fmt.Println("go ping")
	go ping(pings, "passed message")
	fmt.Println("after ping")

	fmt.Println(<-pongs)
	fmt.Println("end1")

	go ping(pings, "passed message2")
	go ping(pings, "passed message3")
	go pong(pings, pongs)
	fmt.Println(<-pings)
	fmt.Println("pings end2")
	fmt.Println(<-pongs)
	fmt.Println("pongs end2")
}
