package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 75}
	// grades := [...]int{97, 85, 75} ->dynamic length
	fmt.Printf("Grades: %v", grades)

	var students [3]string
	fmt.Printf("\nStudents: %v", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("\nStudents: %v", students)
	fmt.Printf("\nStudents #1: %v", students[1])
	fmt.Printf("\nNumber Of Studetns: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// alternative way
	// identityMatrix[0] = [3]int{1, 0, 0}
	// identityMatrix[1] = [3]int{0, 1, 0}
	// identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3, 4}
	b := a
	// b := &a ->pointing to memory of a, what happen to b will affect a
	b[1] = 5
	// reassign array is make whole new array so when changing b elemen, wont affect a
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("array & slices")

	// array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} -> array
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // -> slice
	s1 := array[:]
	s2 := array[3:]
	s3 := array[:6]
	s4 := array[3:6]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	arrayMake := make([]int, 3)
	fmt.Println(arrayMake)

	fmt.Println("append elemen")

	arr_ := []int{}
	fmt.Println(arr_)
	fmt.Printf("Length: %v\n", len(arr_))
	fmt.Printf("Capacity: %v\n", len(arr_))
	arr_ = append(arr_, 1, 2, 3, 4, 5, 6)
	// arr_ = append(arr_, []int{1, 2, 3, 4, 5, 6}...) -> spread operator
	fmt.Println(arr_)
	fmt.Printf("Length: %v\n", len(arr_))
	fmt.Printf("Capacity: %v\n", len(arr_))

	fmt.Println("manipulating elemen // cutting")
	_arr := []int{1, 2, 3, 4, 5}
	fmt.Println(_arr)
	_arr_b := append(_arr[:2], _arr[3:]...)
	fmt.Println(_arr)
	fmt.Println(_arr_b)

	// summary
	// Array
	// collection of items with same type
	// Fixed size
	// declaration style:
	//     - a := [3]int{1,2,3}
	//     - a := [...]int{1,2,3}
	//     - var a [3]int
	// access elemen start from 0 -> zero based index
	//     - a := var a [3]int {1,3,5} -> a[1] = 3
	// len function return size of array
	// copies refer to different underlying data

	// slices
	// backed by array (unfixed size of array)
	// creation style
	//     - slice existing array or slice
	//     - literal style
	//     - using function
	//     -> a := make([]int, 10) -> create slice with capacity and length 10
	//     -> a := make([]int, 10, 100) -> create slice with length = 10 capacity = 100
	// len function return length of slices
	// cap function return length of underlying array
	// append function add elemen to slice -> cause exspensive copy operation if underlying array is too small
	// copies refer to same underlying array

}
