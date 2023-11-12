package main

type printerError struct{}

func (p printerError) Error() string {
	return "cannot access printing device"
}

func main() {
	//panic will call the error of the type it contains, or the string.
	panic(printerError{})

	//an easy way also to send errors.New("cannot access printing device")
	//panic(errors.New("error msg"))
}
