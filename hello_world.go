package main

import "fmt"

func main() {
	test1 := "ABC"
	test2 := "DEF"

	fmt.Println("Hello, World! =>", test1, "/", test2)

	// var student User = User{1, "kim.dev", "1234"}
	student := User{1, "kim.dev", "1234"}

	str_text := fmt.Sprintf("test => %d %s", student.id, student.name)

	fmt.Println("str_text => ", str_text)
}

type User struct {
	id    int
	name  string
	phone string
}
