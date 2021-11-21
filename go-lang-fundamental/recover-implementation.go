package main

import (
	"fmt"
	"log"
)

func main() {
	// Recover is a built-in function that regains control of a panicking goroutine.
	// Recover is only useful inside deferred functions
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	// defer anonymous func
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err) // rethrow panic
		}
	}()
	panic("Something bad happened")
	fmt.Println("done panicking")
}
