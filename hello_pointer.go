package main

import "fmt"

type Actor struct {
	name  string
	hp    int
	speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	actor := Actor{name, hp, speed}

	return &actor
}

func main() {
	actor := NewActor("named", 99, 120)
	fmt.Println(actor.speed)
	fmt.Println(actor.name)
}
