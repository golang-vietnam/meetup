package example

import (
	"fmt"
)

var _Check = func() map[string]int {
	return map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
}

func convertStringToInteger(input string) (int, error) {
	output, ok := _Check()[input]
	if !ok {
		return 0, fmt.Errorf("%s is not satisfy", input)
	}
	return output, nil
}
