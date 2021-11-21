package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := 30

	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too high")
		} else if guess == number {
			fmt.Println("You Got it!")
		}
	}
}
