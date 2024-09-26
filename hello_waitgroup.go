package main

import (
	"fmt"
	"sync"
)

func main() {
	// WaitGroup 생성. 2개의 Go루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(13)

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	// goroutine 안에 또 다른 goroutine 생성
	go func() {
		defer wait.Done() //끝나면 .Done() 호출
		for i := 0; i < 10; i++ {
			go func() {
				defer wait.Done() //끝나면 .Done() 호출
				fmt.Println("i => ", i)
			}()
		}
	}()

	wait.Wait() //Go루틴 모두 끝날 때까지 대기
}
