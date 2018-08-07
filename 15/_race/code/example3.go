package main

import (
	"fmt"
	"time"
)

var count = 0

func incrementJustOnce() {
	if count == 0 {
		time.Sleep(1 * time.Nanosecond)
		count++
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go incrementJustOnce()
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println(count)
}
