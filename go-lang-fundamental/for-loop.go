package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i, j)
		}
	}

	//loop with tag
Loop:
	for i := 0; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i, j)
			if (i * j) > 5 {

				break Loop
			}
		}
	}

}

// summary
// simple loops
// -for initializer; test; incrementer {}
// -for test {}
// -for {}
// exiting loop
// -break
// -continue
// -labels
// loop over collections
// - arrays, slices, maps, strings, channels
// - for k, v := range collections {}
