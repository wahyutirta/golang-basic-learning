package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Hello Go!")
	}()

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
}
