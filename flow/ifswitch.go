package main

import (
	"fmt"
	"math"
)

func main() {
	if true {
		fmt.Println(("The test is true"))
	}

	mapVar := map[string]int{
		"Key1": 10,
		"Key2": 20,
	}
	if pop, ok := mapVar["Key2"]; ok {
		fmt.Println(pop)
	}

	number, guess := 50, 30
	// Go short-circuits the conditionals
	// If the first condition is true, it doesn't evaluate the next OR conditions
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too high")
		} else {
			fmt.Println("You got it!")
		}
	}

	// Another thing to keep in mind with checking equality
	// is when using floating point numbers
	myNum := 0.123
	// This will be false due to floating point approximations
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
	// Preferred:
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	// Flow can also be controlled with switch statements
	switch 3 {
	case 1, 5, 10: // No falling through in Go, but multiple cases
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}
	// Switches can also have initializers
	switch i := 2 + 3; i {
	case 1:
		fmt.Println("Yay")
	default:
		fmt.Println("Yawn")
	}
	// Another syntax, tagless. Allows use of variable inside conditionals
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than equal to ten")
		fallthrough // will fallthrough to the next one
	case i <= 20:
		fmt.Println("Less than equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}
	// Can use interfaces to do type-switching
	var t interface{} = 1
	switch t.(type) {
	case int:
		fmt.Println("t is an int")
		break // can break execution early
		fmt.Println("This will not print")
	case float64:
		fmt.Println("t is a float64")
	case string:
		fmt.Println("t is a string")
	default:
		fmt.Println("t is another type")
	}
}
