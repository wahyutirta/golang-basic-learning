package main

import (
	"fmt"
)

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s)

}

func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func sum_(values ...int) (result int) {
	fmt.Println(values)

	for _, v := range values {
		result += v
	}
	return
}
