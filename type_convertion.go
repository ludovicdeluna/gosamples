//usr/bin/env go run "$0" "$@" ; exit "$?"
package main

import "fmt"

// Untyped const (only for consts)
const my_generic_number = 7       // can be numbers like int, int64, float64...
const my_generic_float = 3.4      // can only be float64
const my_generic_string = "Hello" // can only be strings
// Used into literals, any string or numbers (ever true/false) are constants
func main() {
	var myInt int = 5
	var myFloat float64 = 8.3
	myInt2 := 5 + int(myFloat)
	myInt3 := my_generic_number + myInt            // const used as int
	var myFloat2 float64 = my_generic_number / 2.0 // const used as float64
	var myFloat3 float64 = float64(7) / 2.0        // Other than untyped const => Need a cast
	// Untyped int into a float is OK, but untyped float into an int is not because lost of precision
	// myInt += my_generic_float
	// => Compile error: constant 3.4 truncated to integer
	fmt.Println(myInt, myFloat, myInt2, myInt3, myFloat2, myFloat3)
}
