package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	g.greetPoint()
	fmt.Println(g.name)
}

type greeter struct {
	greeting string
	name     string
}

// methods are the same as functions
// but they are given a context (can be any type, not only structs)
// In this case, a copy of the struct is made
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// If we make a pointer-receiver method, we can operate directly on the values
func (g *greeter) greetPoint() {
	fmt.Println(g.greeting, g.name)
	g.name = "Oops"
}
