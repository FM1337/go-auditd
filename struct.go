package main

import "time"

// Log is our log object.
type Log struct {
	Type   string
	Time   time.Time
	Keys   []string
	Values []string
}
