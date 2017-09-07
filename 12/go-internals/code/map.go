package main

import "fmt"

func main() {
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	for key, value := range commits {
		fmt.Println("Key:", key, "Value:", value)
	}
}
