package main

import (
	"fmt"
	"time"
)

var types map[string]int

// this is used as a quick test to see if I can successfully parse the log types
func countTypes() {
	types = make(map[string]int)

	for _, log := range logs {
		for _, l := range log {
			types[l.Type] = types[l.Type] + 1
		}
	}

	for t, value := range types {
		fmt.Printf("Log type: %s, Count: %d\n", t, value)
	}
}

func testPrint() {
	for _, log := range logs {
		for _, l := range log {
			fmt.Printf("%s: Type: %s\n", l.Time.Format(time.RFC1123), l.Type)
		}
	}
}
