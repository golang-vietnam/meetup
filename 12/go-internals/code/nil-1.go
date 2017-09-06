package main

import (
	"fmt"
	"os"
)

func main() {
	// println(nil) // 1
	println(interface{}(nil)) // 2
	fmt.Println(nil)          // 3
	// println(interface{}(nil) == (*int)(nil)) // 4
	println(nil == error((*os.PathError)(nil))) // 5
}
