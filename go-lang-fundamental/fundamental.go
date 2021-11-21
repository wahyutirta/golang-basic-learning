package main

import (
	"fmt"
	"strconv" // -> int to string conversion module
)

// wring acronym in go full uppoer case
// webHTTP
// webURL

var (
	actorName     string = "Elisabeth Swan"
	companionName string = "Sarah Conor"
	doctorNum     int    = 3
	season        int    = 11
)

var i int = 10

func main() {
	fmt.Printf("%v, %T\n", i, i)

	// go will find the type on variable
	foo := 20
	fmt.Printf("%v, %T\n", foo, foo)
	// shadow variable i, inside main
	var i int = 42

	// cant redeclare variable in smae scope, but you can reassign vaariable
	// i := 13 -> decrale and assign variable
	i = 43

	var j float32 = 27

	k := 99

	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

	// variable type convertion
	var num int = 42
	fmt.Printf("%v, %T\n", num, num)

	var numFloat float32
	numFloat = float32(num)
	fmt.Printf("%v, %T\n", numFloat, numFloat)

	var numString string
	numString = strconv.Itoa(num)
	fmt.Printf("%v, %T\n", numString, numString)
}
