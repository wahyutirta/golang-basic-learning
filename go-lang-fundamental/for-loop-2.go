package main

import (
	"fmt"
)

func main() {
	statePopulation := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	s := []int{1, 2, 3, 4}
	for k, v := range statePopulation {
		fmt.Println(k, v)
	}
	for k, v := range s {
		fmt.Println(k, v)
	}

	aString := "hello go"
	for k, v := range aString {
		fmt.Println(k, string(v))
	}

	// iterate value  only
	for _, v := range aString {
		fmt.Println(string(v))
	}
}
