//usr/bin/env go run "$0" "$@" ; exit "$?"
package main

import "fmt"

func main() {
	var str_msg string = "Hello, 世界"
	var utf_msg = []rune(str_msg)
	var str_idx_max int

	fmt.Println(str_msg)
	fmt.Println("----------")
	for utf_idx, str_idx := 0, 0; utf_idx < len(utf_msg); utf_idx += 1 {
		if byte(str_msg[str_idx]) > 128 {
			str_idx_max = str_idx + 3
		} else {
			str_idx_max = str_idx + 1
		}
		fmt.Printf(
			"%s => %d\n",
			string(utf_msg[utf_idx]),
			[]byte(str_msg[str_idx:str_idx_max]),
		)
		str_idx = str_idx_max
	}
	fmt.Println("----------")
	fmt.Printf("%d => %d\n", len(utf_msg), len(str_msg))
}
