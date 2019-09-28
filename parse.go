package main

import (
	"strconv"
	"strings"
)

// logs is our map containing the logs
var logs map[string][]log

func getLogID(msg string) string {
	return strings.TrimLeft(strings.TrimRight(msg, "):"), "msg=audit(")
}

func getTimestamp(msg string) int64 {
	id := getLogID(msg)
	t := strings.Split(id, ":")[0]
	ts, err := strconv.ParseInt(strings.Split(t, ".")[0], 10, 64)
	if err != nil {
		panic(err)
	}
	return ts
}
