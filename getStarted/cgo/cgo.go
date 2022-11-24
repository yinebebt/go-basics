// cgo enables go package to import/use data from c files and vise versa.
// once we import C, cgo will look docs of c equivalent codes and will parse and use it.
// C is a “pseudo-package”, a special name interpreted by cgo as a reference to C’s name space
package main

// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
//		return f();
// }
//
// int fortytwo()
// {
//	    return 42;
// }
import "C"
import "fmt"

func main() {
	f := C.intFunc(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))
	// Output: 42
}
