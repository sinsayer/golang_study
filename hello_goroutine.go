package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	done := make(chan bool)

// 	values := []string{"a", "b", "c"}
// 	for _, v := range values {
// 		go func() {
// 			fmt.Println(v)
// 			done <- true
// 		}()
// 	}

// 	// wait for all goroutines to complete before exiting
// 	for _ = range values {
// 		<-done
// 	}
// }
