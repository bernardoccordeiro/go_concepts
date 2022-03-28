package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// In most languages, OS threads are used.
// That means they have individual function call stacks,
// dedicated to executing the code handed to that thread.
// These can be very large and take time for the app to set up.
// So, ideally, they need to be used conservatively. Since
// creating and destroying threads can be expensive.

// In Go, a different model is used. A "green thread" is used instead.
// It is a lightweight thread that an internal scheduler in Go assigns to OS threads, taking turns with every CPU thread that is available, assigning the Go routines some processing time on those threads.
// The developer doesn't have to interact directly with the OS threads, but with the high-level goroutines.
// The advantage is that goroutines are very lightweight, cheap to create and destroy.

// WaitGroups are meant to sycnrhonize multiple Goroutines
var wg = sync.WaitGroup{}
var counter = 0

// A Mutex is basically a Lock that the app will honor
var m = sync.RWMutex{}

func main() {
	var msg = "Hello"
	// passing a variable to a goroutine avoids
	// undesired race conditions
	wg.Add(1) // Adds 1 to WaitGroup counter
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() // Will decrement 1 from the WaitGroup
	}(msg)
	msg = "Goodbye"
	wg.Wait() // When the counter reaches 0 again, it will stop waiting

	// With race condition, to show compiler option
	// Using go run -race file.go
	var msg_race = "Hi"
	go func() {
		fmt.Println(msg_race)
	}()
	msg_race = "Goodbye"
	time.Sleep(100 * time.Millisecond)

	// How about with multiple goroutines?
	runtime.GOMAXPROCS(100) // Will set the number of OS threads
	// Could be a good tweaking parameters before production

	// This application actually has no paralellism
	// But is good to show how locks work
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

// Best Practices
// - Don't create goroutines in libraries - let consumer control concurrency
// - When creating a goroutine, know how it will end - avoid subtle memory leaks
// - Check for race conditions at compile time
