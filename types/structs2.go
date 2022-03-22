package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"` // use a tag
	Origin string
}

// Embedding a struct in another struct
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

// Embedding *can* transfer behavior to other structs
// However, it is not interchangeable back to the "parent" struct
// It is not object-oriented programming inheritance
// For this use case, Golang's interfaces might be considered instead

func main() {
	// Literal syntax, you need to know the internal structure:
	// b := Bird{
	//		Animal: Animal{Name: "Emu", Origin: "Australia"},
	//		SpeedKPH: 48,
	//		CanFly: false,
	// }

	// Creating an object and manipulating from the outside
	// provides more abstraction
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b.Name)

	// Golang can use tags to provide additional information
	// about a struct's fields
	// They work like a docstring of sorts
	// They can be extracted for use by other libraries, using the reflect package
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
