package main

import (
	"fmt"
)

func main() {
	sayMessage("Hello Go", 2)

	s := sum(1, 2, 3, 4, 5)
	fmt.Println(s)

	sP := sumPointer(1, 2, 3, 4)
	fmt.Println(*sP)

	d, err := divide(5.0, 1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// Anonymous function
	func() {
		fmt.Println("Hello Go!")
	}() // immediately execute it
	f := func() {
		fmt.Println("Attribute it to variable")
	}
	f()
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

// No need to declare the type of all of them in this case
func sayMessageSameTypeParameters(msg1, msg2 string) {
	fmt.Println(msg1, msg2)
}

// Using variadic parameters
// Must be the last one in the list of parameters
func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// In this case, Go promotes the returned value from local memory (stack) to shared memory (heap)
func sumPointer(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

// Name and declare the return variable beforehand
// Doesn't need to redeclare it or say what is returned
func sumNamedReturn(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

// Return an error value if something went wrong
// Very idiomatic with Golang
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
