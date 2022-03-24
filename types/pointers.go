package main

import (
	"fmt"
)

func pointerArithmeticTest() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	// c := &a[1] - 8
	// Normally this would give the address of the previous element
	// However, Go does not allow pointer arithmetic
	// IF NECESSARY, there is an "unsafe" package that allows for this
	fmt.Printf("%v %p %p\n", a, b, c)
}

type myStruct struct {
	foo int
}

func structPointerTest() {
	var ms *myStruct
	fmt.Println() // uninitialized pointers are nil pointers
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	ms = new(myStruct)
	(*ms).foo = 42 // parentheses to indicate precedence
	ms.foo = 42    // same as above, but reads more cleanly, syntactic sugar
	fmt.Println(ms)
}

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b, *b) // * is the dereferencing operator
	a = 27
	fmt.Println(a, b, *b)

	fmt.Println("Testing pointer arithmetic")
	pointerArithmeticTest()

	fmt.Println("Testing with structs")
	structPointerTest()
}
