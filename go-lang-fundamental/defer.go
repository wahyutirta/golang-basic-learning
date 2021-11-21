package main

import (
	"fmt"
)

func main() {
	// defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns. In other words,
	// defer function or method call arguments evaluate instantly, but they don't execute until the nearby functions returns
	// defer runs multiple call defer calls by stack, last in first out
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")

}
