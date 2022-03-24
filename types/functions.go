package main

import (
	"fmt"
)

func main() {
	sayMessage("Hello Go", 2)
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

// No need to declare the type of all of them in this case
func sayMessageSameTypeParameters(msg1, msg2 string) {
	fmt.Println(msg1, msg2)
}
