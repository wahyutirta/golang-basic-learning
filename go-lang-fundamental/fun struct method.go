package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "go",
	}
	g.greet()

	j := greeter{
		greeting: "Hello",
		name:     "java",
	}
	j.greet()
}

type greeter struct {
	greeting string
	name     string
}

//this func refer to memory ob struct g
func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
