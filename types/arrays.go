package main

import (
	"fmt"
)

func main() {
	// Need to specify the size of the array
	// Are contiguous in memory
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	// can also declare arrays with ... syntax
	// which means "create an array big enough to hold this data"
	more_grades := [...]int{100, 90, 80, 70}
	fmt.Printf("More grades: %v\n", more_grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Maximum number of students: %v\n", len(students))

	// IMPORTANT: In Golang, arrays are actually values
	// This means, assigning a variable to an existing array
	// actually creates a copy of it
	fmt.Printf("Arrays\n")
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b) // will produce different results
	// If using the same array is required, use pointers

	// Golang also has what are called 'slices'
	// They are similar to arrays

	fmt.Printf("Slices\n")
	a1 := []int{1, 2, 3} //slice - just brackets with no number
	b1 := a1             // slices are actually passed by reference
	c := a1[1:]          // similar to Python, but reference same array
	b1[1] = 5
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c)
	fmt.Printf("Length: %v; Capacity: %v\n", len(a1), cap(a1))
	// Slices also have what is called a capacity
	// which is the size of the underlying array

	// We can use the make function to create a slice and set its capacity
	a2 := make([]int, 3, 100)
	fmt.Println(a2)
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))

	// Slices can be appended to
	a3 := []int{}
	fmt.Printf("Slice: %v; Length: %v; Capacity: %v\n", a3, len(a3), cap(a3))
	a3 = append(a3, 1)
	fmt.Printf("Slice: %v; Length: %v; Capacity: %v\n", a3, len(a3), cap(a3))
	a3 = append(a3, []int{2, 3, 4, 5}...) // use the ... operator to append two slices
	fmt.Printf("Slice: %v; Length: %v; Capacity: %v\n", a3, len(a3), cap(a3))

	// If you want to use stack operations and pop values
	// Use a combination of slices and append
	st := []int{1, 2, 3, 4, 5}
	fmt.Println(st)
	// st2 := st[1:] - pop first element
	// st2 := st[:len(st) - 1] - pop last element
	st2 := append(st[:2], st[3:]...) // remove 3rd element
	fmt.Println(st2)
	fmt.Println(st) // CAUTION: The underlying array might get messed up
}
