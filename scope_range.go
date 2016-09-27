//usr/bin/env go run "$0" "$@" ; exit "$?"
package main

import "fmt"

func iterateOnPointer() <-chan *int {
	done := make(chan *int)
	go func() {
		defer close(done)
		for i := 0; i < 6; i++ {
			fmt.Printf("(%d)", i)
			done <- &i // Return memory address of 'i' which content change on each loop
		}
	}()
	return done
}

func iterateOnValue() <-chan int {
	done := make(chan int)
	go func() {
		defer close(done)
		for i := 0; i < 6; i++ {
			fmt.Printf("(%d)", i)
			done <- i // Return copy of 'i' content, the safe way to use channel
		}
	}()
	return done
}

func main() {
	// Unsafe
	fmt.Print("Pointer : ")
	for v := range iterateOnPointer() {
		fmt.Printf("[%d]", *v) // => (0)(1)[1][1](2)(3)[3][3](4)(5)[5][5]
		// Returned                         ^  ^        ^  ^        ^  ^
	}
	// Safe
	fmt.Print("\nValue :   ")
	for v := range iterateOnValue() {
		fmt.Printf("[%d]", v) // =>  (0)(1)[0][1](2)(3)[2][3](4)(5)[4][5]
		// Returned                         ^  ^        ^  ^        ^  ^
	}
  // Unsafe with other Go Routine
  fmt.Print("\nGo on Pointer : ")
  for v := range iterateOnPointer() {
    go func() {
      for i := 0 ; i < 5 ; i++ {
      }
      fmt.Printf("[%d]", *v)
    }()
  }
  // With or without a value, use in another Go need also a copy.
  // Bad way :
  fmt.Print("\nWithout scope (bad) : ")
  for v := range iterateOnValue() {
    go func() {
      for i := 0 ; i < 5 ; i++ {
      }
      fmt.Printf("[%d]", v) // (0)(1)[1][1](2)(3)[3][3](4)(5)[5][5]
      // Results                      ^  ^        ^  ^        ^  ^
    }()
  }
  // Good way :
  fmt.Print("\nWith scope (good) : ")
  for v := range iterateOnValue() {
    go func(value int) {
      for i := 0 ; i < 5 ; i++ {
      }
      fmt.Printf("[%d]", value) // (0)(1)[1][0](2)(3)[3][2](4)(5)[5][4]
      // Results                          ^  ^        ^  ^        ^  ^
    }(v)
  }
}
