//usr/bin/env go run "$0" "$@" ; exit "$?"
package main

import "fmt"

func main() {
	AssignUnderlyingType()
}

func AssignUnderlyingType() {
	type Dictionary map[string]int
	type Scores map[string]int
	var myDictionary Dictionary

	// Now we can do:
	myDictionary = Dictionary{"Ludo": 5}         // Use the variable's type
	myDictionary = map[string]int{"Ludo": 5}     // Use the underlying variable's type
	myDictionary = Dictionary(Scores{"ludo": 5}) // Extract the underlying type of Score

	// We can't do:
	// myDictionary = Scores{"ludo": 5}
	// => Compile Error:
	//    cannot use Scores literal (type Scores) as type Dictionary in assignment
	fmt.Printf(
		"myDictionary is a %#v and contain a %v\n",
		myDictionary,
		myDictionary,
	)
	// => myDictionary is a main.Dictionary{} and contain a map[Ludo: 5]
}
