package main

import (
	"fmt"
	"time"
)

var count = 0

func counter() {
	for j := 0; j < 100; j++ {
		count++
	}
}

func main() {
	for i := 0; i < 100; i++ {
		go counter()
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println(count)
}
