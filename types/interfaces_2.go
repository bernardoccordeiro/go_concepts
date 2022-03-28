package main

import (
	"fmt"
)

// if all of the methods of the type you're implementing
// are value receivers, then you can implement the interface by
// using the value of the type
// this is because the method set of a value type is the value receiver methods
// else, if at least one of the methods is a pointer receiver
// then the interface needs to be implemented by a pointer to the type
// a pointer to the type also works in case all of the methods are value receivers
// this happens because the method set of a pointer type is the pointer receiver methods + the value receiver methods
func main() {
	var wc WriterCloser = &myWriterCloser{}
	fmt.Println(wc)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct{}

func (mwc *myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil
}

// Interface best practices
// - Prefer many small interfaces instead of large monolithic ones; examples - io.Writer, io.Reader, interface{}
// - Don't export interfaces for types that will be consumed
// - Do export interfaces for types that will be used by the package
// - Design functions and methods to receive interfaces whenever possible
