package main

import (
	"fmt"
)

func main() {

	statePopulation := make(map[string]int)

	statePopulation = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	fmt.Println(statePopulation)

	// fmt.Println(statePopulation["Ohio"])
	statePopulation["Georgia"] = 10310371
	fmt.Println(statePopulation)
	delete(statePopulation, "Georgia") // delete map elemen
	fmt.Println(statePopulation)
	// m := map[[3]int]string{}
	// fmt.Println(statePopulation)

	value, ok := statePopulation["Ohio"] // cek value, keberadaan key -> true artinya ada -> false artinya key tidak ada
	fmt.Println(value, ok)
	fmt.Println(len(statePopulation))
}
