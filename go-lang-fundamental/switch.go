package main

import (
	"fmt"
)

func main() {
	num := 3
	switch num {
	case 1, 5, 10:
		fmt.Println("One, Five or Ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	//switch with operator
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("One, Five or Ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	// fallthrough -> run all cases
	aNum := 10
	switch {
	case aNum <= 10:
		fmt.Println("less or equal to ten")
		fallthrough //
	case aNum >= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	var aVar interface{} = 1.0
	switch aVar.(type) {
	case int:
		fmt.Println("it is a int")
	case float64:
		fmt.Println("it is a float64")
	case string:
		fmt.Println("it is a string")
	default:
		fmt.Println("it is another type")
	}
}
