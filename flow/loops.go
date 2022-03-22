package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// two-variable loop
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	// the initializers are scoped to the loop
	// so this would not run:
	// fmt.Println(i)

	// To be able to use an index value outside loop, declare it beforehand
	var i = 2
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Printf("i still exists and is %v\n", i)

	// for can be used as while
	// semi-colons may be removed in this case
	i = 1
	for i < 5 {
		fmt.Println(i)
		i++
	}
	// it may also be left empty and used with break
	for {
		i++
		break
	}
	// another flow control keyword is continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	// Golang supports Labels, useful for breaking nested loops
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
	// Looping through collections
	// Works *kinda* similar to Python's enumerate or
	// items() function for Python dicts
	// Can use _ to ignore one of the outputs
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
	st := "hello Go!"
	for k, v := range st {
		fmt.Println(k, string(v))
	}
	mp := map[string]int{"K1": 10, "K2": 20}
	for k, v := range mp {
		fmt.Println(k, v)
	}
}
