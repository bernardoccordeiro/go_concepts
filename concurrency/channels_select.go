package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // no data allocation, but can send the info that a message has been sent
// preferable to make(chan bool), which has to allocate memory

func main() {
	go logger()
	// The defer function can help here to draw closing strategies
	// defer func() {
	// 	close(logCh)
	// }()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	// Goroutines are being forcibly torn down at the end
	// But using an empty struct channel, you can control that better
	doneCh <- struct{}{}
}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
			// if you add default, it becomes a non-blocking select statement
			// and if no messages are ready in the channel, it will execute the default
			// default:
		}
		// could also loop the channel and let it forcibly close
		// but could leave memory leaks
		// entry := range logCh {
		// fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
}
