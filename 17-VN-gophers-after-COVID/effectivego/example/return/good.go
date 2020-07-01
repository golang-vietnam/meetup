package example

func handleSlice(input []string) {
	for i := range input {
		if input[i] == "" {
			// handle if needed
			continue
		}

		// do some stuff
	}
}
