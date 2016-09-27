//usr/bin/env go run "$0" "$@" ; exit "$?"
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Message struct {
	Name string `json:"name,omitempty"` //Flags: without space. omitempty
	Body string `json:"body,omitempty"` //will suppress the entry if
	Time int64  `json:"time,omitempty"` //this is the default value
}

func main() {
	now := time.Now()
	msg := Message{"Alice", "Hello World ñ !", int64(now.Unix())}
	fmt.Printf("Initial message : %#v\n", msg)
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		os.Exit(1)
	}
	jsonStr := string(jsonMsg)
	fmt.Printf("Resulting JSON (string from a []byte) : %s\n", jsonStr)
	// Decode
	var msgAgain Message
	if err := json.Unmarshal([]byte(jsonStr), &msgAgain); err != nil {
		os.Exit(1)
	}
	fmt.Printf("Decoding JSON (go back to Message) : %#v\n", msgAgain)
	// Partial Decode
	var otherStr = "{\"name\":\"Alice\",\"body\":\"Hello World ñ !\"}"
	fmt.Printf("Other JSON (string) = %#v\n", otherStr)
	var msgOther Message
	if err := json.Unmarshal([]byte(otherStr), &msgOther); err != nil {
		os.Exit(1)
	}
	fmt.Printf("Partial JSON decoding (go back to Message) : %#v\n", msgOther)
}
