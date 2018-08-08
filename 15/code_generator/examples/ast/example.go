package main

func A() (int, error) {
	err := B()

	return 10, nil
}

func B() error {
	err := C()

	return nil
}

func C() error {
	return nil
}
