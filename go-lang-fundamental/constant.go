package main

import (
	"fmt"
)

const myConst int16 = 10

const (
	a = iota // == 0
	b = iota // == 1
	c = iota // == 2
)

const (
	//iota automatically reset and find pattern inside
	_  = iota //dump zero value
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeefinancials

	canSeeAfrica
	canSeeAsian
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	const myConst int = 42

	fmt.Printf("%v, %T\n", myConst, myConst)

	// cant assign value to const when the value  has to be determine at runtime (exe)
	// const mySine float64 = math.Sin(1.57)
	// we can shadow const valiable in GO

	const a = 42
	var b int16 = 27
	fmt.Printf("%v, %T\n", a+b, a+b)
	// if there is no spesific data type, compiler see a as 42 then we cann add a + b like 42 + b

	// iota app example
	fmt.Println("iota memory converter")
	fSize := 4000000000.
	fmt.Printf("%.2fGB", fSize/GB)

	var roles byte = isAdmin | canSeefinancials | canSeeEurope //only take the true value by or
	fmt.Println("\nroles in byte")
	fmt.Printf("%b\n", roles)

	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isAdmin)

	// summary
	//  const is immutable, but can be shadowed
	// replaced by the compiler at the compiler time, value must be calculable at compiler time not while running

	// named like variables
	// PascalCase for exported constants
	// camelCase for internal constants

	//  Typed constants work like immutable variable
	//  cam interoperate only with same type

	// Untyped constants work life literals
	//  can interoperate with similar types\

	// enumerated constants
	// speciaal symbol iota allows related constants tobe craeted easily
	// itoa starts at 0 in each const block and increments by one
	// watch out of constants values that match zero values for variable

	// enumerated expressions
	// operations that can be detemined at compile time are allowed
	// - arithmath
	// - bitwise operations
	// - bishifting
}
