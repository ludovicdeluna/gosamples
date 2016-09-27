//usr/bin/env go run "$0" "$@" ; exit "$?"
package main

// From question :
// https://forum.golangbridge.org/t/using-a-switch-to-handle-different-types-in-an-object-map/3351/6
// Original question
/*
I want to print on a web page the data for the userSessionID but the data is of
different types. For example I tried to do a .(string) type assertion for Value
in the Object Map but I got this error:

    panic...interface conversion: interface {} is []uint8, not *string
*/

// Interesting Answer on Stack Overflow (on related subject)
/*
I need to read [100]byte to transfer a bunch of string data.
Because not all the string is precisely 100 long, the remaining part of the byte array are padded with 0s.
If I tansfer [100]byte to string by: string(byteArray[:]), the tailing 0s are displayed as ^@^@s.
In C the string will terminate upon 0, so I wonder what's the best way of smartly transfer byte array to string.
*/

import (
	"bytes"
	"fmt"
)

// Playground here to reproduce panic
func will_panic() {
	var x uint8 = 8
	var value interface{}
	value = x
	str := value.(*string) // When forcing to use another type, it fail
	fmt.Println(str)
}

func main() {
	without_0_byte_escape()
	fmt.Println("----------------------")
	with_0_byte_escap()

}

// Use sessions from this package :
// http://www.gorillatoolkit.org/pkg/sessions
// sessions.Val return []uint8 (identical to []byte)
// Solution : Convert []byte to string
// See: http://stackoverflow.com/questions/14230145/what-is-the-best-way-to-convert-byte-array-to-string
//
// This is the Right answer if you don't pad a prefixed []byte (as requested initially)
func without_0_byte_escape() {
	fmt.Println("Without Escape")
	var byteArray [5]byte
	byteArray[0] = 65

	// This method is short, but resulting can contains ^@ characters (byte 0)
	fmt.Printf("Content -> %v\n", byteArray[:])
	fmt.Printf("String -> %#v\n", string(byteArray[:])) // Println will escape bytes of 0, but it's not the case for all object.
}

// Because 0 byte are not escaped (array of 5 slots), we have to find where
// are 0. Here possible solution :
func with_0_byte_escap() {
	fmt.Println("With Escape")
	var byteArray [5]byte
	byteArray[0] = 65

	n := bytes.IndexByte(byteArray[:], 0)
	s := string(byteArray[:n])
	fmt.Printf("Content -> %v\n", byteArray[:])
	fmt.Printf("String -> %#v\n", s)
}
