package main

import (
	"fmt"
)

func main() {
	sayGreeting("Hello", "Stacey")
	sayGreetings("Hello", "Stacey")
}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

func sayGreetings(greeting string, name string) {
	fmt.Println(greeting, name)
}
