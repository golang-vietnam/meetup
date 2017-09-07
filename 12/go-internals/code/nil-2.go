package main

import "fmt"

func main() {
	fmt.Printf("Func type nil:%#v\n", (func())(nil))
	fmt.Printf("Map type nil:%#v\n", map[string]string(nil))
	fmt.Printf("Slice type nil:%#v\n", []string(nil))
	fmt.Printf("Interface{} type nil:%#v\n", nil)
	fmt.Printf("Channel type nil:%#v\n", (chan struct{})(nil))
	fmt.Printf("Pointer type nil:%#v\n", (*struct{})(nil))
	fmt.Printf("Pointer type nil:%#v\n", (*int)(nil))
}
