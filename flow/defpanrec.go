package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func deferTest() {
	// defer executes at the end
	// but evaluates parameters when it's called
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func multipleDeferTest() {
	// if you have several defers, they execute in reverse
	// order, basically LIFO order
	defer fmt.Println(1)
	defer fmt.Println(2)
}

// panics are like raise in Python
// In Go, a lot of cases that would raise exceptions in other languages
// Just raise an error flag in Go, and the developer chooses
// If he'll raise a panic or just keep going
func panicAndDeferTest() {
	// panics occur after deferred function is executed
	fmt.Println("start")
	defer fmt.Println("This was deferred")
	panic("something bad happened")
	fmt.Println("end")
}

func panicAndRecoverTest() {
	// recover is only useful inside a deferred function
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// if an error should continue, put another panic here
		}
	}()
	panic("Something bad happened")
	fmt.Println("Done panicking")
}

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// defer can be useful to open a resource, and remember to close it
	// but defer that to the end of the function
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	deferTest()
	multipleDeferTest()

	fmt.Println("Start Panic Recover test")
	panicAndRecoverTest()
	fmt.Println("Recovered from panic in main function")

	panicAndDeferTest()
}
