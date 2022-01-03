package main

import (
	"fmt"
	"strconv"
)

// Goland is a strongly typed language. Variables have a single type that needs to be known at compile time.

// Variables can be declared in a variety of ways, which will be shown in this module.

// For instance, you may want to declare a package-level variable, which is accessible throughout the package.
// Here, you cannot use the := syntax, only the other variable declaration syntaxes.
var m int = 100

func main() {
	// When declaring a variable, you can specify the type of the variable.
	var i int
	i = 42
	var j float32 = 10

	// The compiler may also infer the type of the variable based on the value assigned to it.
	var k = 20.

	// This approach, however, will limit the types that can be assigned to the variable.
	// For instance, k will have a type float64 in this case - which might not be what you want, if instead you were aiming for a float32.
	fmt.Printf("k value: %v, k type: %T\n", k, k)

	// Variables might also be declared and initialized directly using the := syntax
	l := 70

	fmt.Println(i, j, k, l, m)

	// It is also possible to declare a block of variables at once.
	// This is also useful to group different variables into a single context, and make readability better.
	var (
		actorName    string = "Elisabeth Sladen"
		companion    string = "Sarah Jane Smith"
		doctorNumber int    = 3
		season       int    = 11
	)
	fmt.Println(actorName, companion, doctorNumber, season)

	// Variable declaration may also suffer what is called "shadowing" - which is when a variable is declared in a scope that is already using the same name.
	// For instance, we are able to redeclare the variable m.
	fmt.Printf("Module-level m is %v\n", m)
	var m float32 = 150.0
	fmt.Printf("Function-level m is %v\n", m)

	// One feature Goland brings is that ALL declared variables need to be used.
	// Uncommenting the following line will result in a compiler error:
	// var unusedVariable string = "This variable is never used"

	// Variables have 3 scopes: block scope; package scope; and global scope.
	// Block scope is the scope of a function, method, or closure.
	// Package scope is the scope of a package. Any file that is part of the package may see a package-level variable. They start with lowercase letters (e.g.: var m int)
	// Global scope variables can be seen from anywhere, including other modules that use the package. To declare this scope, a package-level variable needs to start with an uppercase letter (e.g.: var M int).

	// Type conversion
	var a int = 10
	fmt.Printf("%v, %T\n", a, a)

	var b float32
	b = float32(a)
	fmt.Printf("%v, %T\n", b, b)
	//When working with strings:
	var c int = 42
	var d string
	d = strconv.Itoa(c)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
}
