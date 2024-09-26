package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	fmt.Println("start")
	done := make(chan bool, 1)
	fmt.Println("before worker")
	go worker(done)
	fmt.Println("after worker")

	<-done
	fmt.Println("end")
}
