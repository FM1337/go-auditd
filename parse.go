package main

import "strings"

// logs is our map containing the logs
var logs map[string][]log

func getLogID(msg string) string {
	return strings.TrimLeft(strings.TrimRight(msg, ")"), "(")
}
