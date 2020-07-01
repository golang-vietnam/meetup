package example

func main() {
	var i interface{} = "hello"

	f, ok = i.(float64)
	if !ok {
		// don't use f
	}
}
