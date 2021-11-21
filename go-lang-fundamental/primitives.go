package main

import (
	"fmt"
)

func main() {
	n := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", n, n)

	// secara default golang akan mendefinisikan suatu variabel boolean bernilai 0
	var booleanVar bool
	fmt.Printf("%v, %T\n", booleanVar, booleanVar)

	a := 10 // 1010
	b := 3  // 0011
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// a := 10 1010
	// b := 3  0011
	// a and b = 0010
	// a or b = 1011
	// a xor b 0 1010
	// a and not 0100

	fmt.Println("\nbinary operator")
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)

	fmt.Println("\nbit shifting")
	// bit shifting
	c := 8
	fmt.Println(c << 3) //geser bit 3 langkah ke kiri (dikali)
	fmt.Println(c >> 3) //geser bit 3 langkah ke kanan (dibagi)

	//complex numbers
	fmt.Println("\nComplex Number")
	var numCom complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", real(numCom), real(numCom))
	fmt.Printf("%v, %T\n", imag(numCom), imag(numCom))
	//bigger complex numbers
	var bigCom complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", bigCom, bigCom)

	//string
	fmt.Println("\nString Lateral")
	s := "this is a string literal"
	// -> array, collection of letters, but string storred as int8
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))

	s2 := " this is also a second string literal"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	byteString := []byte(s)
	fmt.Printf("%v, %T\n", byteString, byteString)

}
