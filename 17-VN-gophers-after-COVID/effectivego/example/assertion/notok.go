package example

import "fmt"

func main() {
	var i interface{} = "hello"

	f = i.(float64) // run time panic: interface conversion: interface {} is string, not float64
	fmt.Println(f)
}
