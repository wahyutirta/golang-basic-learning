package main

import (
	"fmt"
	"math"
)

func main() {
	myNum := 0.1

	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}
