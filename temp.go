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

	// for t, value := range types {
	// 	fmt.Printf("Log type: %s, Count: %d\n", t, value)
	// }
}

func testPrint() []string {
	tmpLogs := []string{}
	for _, key := range keys {
		mappedLog := logs[key]
		for _, l := range mappedLog {
			tmpString := ""
			tmpString += fmt.Sprintf("%s: Type: %s, ID: %s, ", l.Time.Format(time.RFC1123), l.Type, l.ID)
			for i, k := range l.Keys {
				if k == "proctitle" {
					tmpString += fmt.Sprintf("%s: %s", k, decryptProctitle(l.Values[i]))
				} else {
					tmpString += fmt.Sprintf("%s: %s", k, l.Values[i])
				}
				tmpString += "\n"
			}
			tmpLogs = append(tmpLogs, tmpString)
		}
	}
	return tmpLogs
}
