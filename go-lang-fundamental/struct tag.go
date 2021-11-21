package main

import (
	"fmt"
	"reflect" // tag library
)

type Animal struct {
	Name   string `required max:100` //tag seperated with backtick
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag) // tag nantinya akan diproses oleh code tersendiri
}
