package main

import (
	"fmt"
)

func main() {
	statePopulations := make(map[string]int, 30)
	statePopulations = map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612539,
	}
	pop, ok := statePopulations["Florida"]
	fmt.Println(pop, ok)
	fmt.Println(len(statePopulations))

	sp := statePopulations // passes by reference
	delete(sp, "Texas")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}
