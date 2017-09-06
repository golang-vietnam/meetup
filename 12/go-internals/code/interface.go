package main

import "fmt"

type I interface {
	M() string
}
type T struct {
	name string
}

func (t T) M() string {
	return t.name
}
func Hello(i I) {
	fmt.Printf("Hi, my name is %s\n", i.M())
}
func main() {
	Hello(T{name: "Jon Snow"})
}
