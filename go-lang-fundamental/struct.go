package main

import "fmt"

type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

func main() {
	aDoctor := Doctor{
		Number:    3,
		ActorName: "Jhon Pert",
		Companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane",
		},
	}
	fmt.Println(aDoctor)
	aDoc := struct{ name string }{name: "Jhon Sena"} // anonymous struct struct{key type}{key: value}
	anotherDoc := aDoc                               // refering to another whole new data
	// anotherDoc := &aDoc                               // refering to aDoc data
	anotherDoc.name = "Tom Baker"
	fmt.Println("Anonimous Struct")
	fmt.Println(aDoc)
	fmt.Println(anotherDoc)

	// structs
	// 	collections of disparate data type that describe a single concept
	// keyed by named fields
	// normally created as types, but anonymous structs are allowed
	// structs are value types
	// no inheritance, but can use composition via embedding
	//
}
