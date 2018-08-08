package main

import "fmt"
import "./enum"

func main() {
	var c enum.ChangeType = enum.Rename

	fmt.Println(c.String())
}
