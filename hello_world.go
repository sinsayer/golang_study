package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var student User = User{1, "kim.dev", "1234"}

	str_text := fmt.Sprintf("test => %d %s", student.id, student.name)

	fmt.Println("str_text => ", str_text)
}

type User struct {
	id    int
	name  string
	phone string
}
