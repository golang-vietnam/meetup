package example

func handleSlice(input []string) {
	for i := range input {
		if input[i] != "" {
			// do some stuff
		} else {
			// handle zero value of element
		}
	}
}
