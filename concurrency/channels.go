package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // declare the data type that will flow through it
	wg.Add(2)
	go func(ch <-chan int) { //receive-only channel
		i := <-ch
		fmt.Println(i)
		// ch <- 27 : this would fail since it is receive-only
		wg.Done()
	}(ch)
	go func(ch chan<- int) { //send-only channel
		i := 42
		ch <- i
		i = 27 // doesn't affect channel, since it passes by value
		// k := <- ch : this would fail since it is send-only
		wg.Done()
	}(ch)
	wg.Wait()

	ch2 := make(chan int, 50) // buffered channel
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		// syntactically the same, but allows for me flexibility
		// for {
		// 	if i, ok := <- ch; ok {
		// 		fmt.Println(i)
		// 	} else {
		// 		break
		// 	}
		// }
		wg.Done()
	}(ch2)
	go func(ch chan<- int) {
		ch2 <- 42
		ch2 <- 27
		close(ch) // prevents deadlock in the receiver function
		wg.Done()
	}(ch2)
	wg.Wait()
}
