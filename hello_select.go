package main

import (
	"fmt"
)

func main() {
	// c1 := make(chan string)
	// c2 := make(chan string)

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	c1 <- "one"
	// }()
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	c2 <- "two"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println("received", msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println("received", msg2)
	// 	}
	// }

	/**
	timeout : After 만큼 기다리고 끝나면 종료
	*/
	// c1 := make(chan string, 1)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	c1 <- "result 1"
	// }()

	// select {
	// case res := <-c1:
	// 	fmt.Println(res)
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("timeout 1")
	// }

	// select 에서 default 를 사용하면 데이터가 없어도 계속 기다리지 않고 빠져나옴.
	c1 := make(chan string)
	go func() {
		c1 <- "result 1"
	}()

	// for i := 0; i < 2; i++ {
	select {
	case res := <-c1:
		fmt.Println(res)
	default:
		fmt.Println("no data in c1")
	} //non-blocking 형태의 receive
	select {
	case c1 <- "msg":
		fmt.Println("msg sent")
	default:
		fmt.Println("no receiver ready")
	} //non-blocking 형태의 sender

	select {
	case res := <-c1:
		fmt.Println(res)
	default:
		fmt.Println("no data in c2")
	} //non-blocking 형태의 receive

}
