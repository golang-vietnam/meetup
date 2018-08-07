package main

import (
	"fmt"
	"time"
)

func main() {
	var message string

	go func() {
		message = "Hello, Go"
	}()
	time.Sleep(4 * time.Microsecond)
	fmt.Println(message)
}
