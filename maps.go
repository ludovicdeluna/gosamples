//usr/bin/env go run "$0" "$@" ; exit "$?"
package main

import "fmt"

type Scores map[string]int
type Scorings map[string]*int

func main() {
	scores := Scores{"Ludo": 5}
	fmt.Println("Scores is a map: Dictionary of keys => value as a special pointer")
	fmt.Printf("Scores is now : %v\n\n", scores)
	fmt.Println("I add a new score (without pointer)")
	addScore(scores, "Thomas", 4)
	fmt.Printf("Scores is now : %v\n\n", scores)

	var points [4]int
	points[0] = 5
	scorings := Scorings{"Ludo": &points[0]}
	fmt.Printf("socrings is : %#v\n", scorings)
}

func addScore(s Scores, name string, points int) {
	s[name] = points
}
