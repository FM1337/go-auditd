package main

import "time"

// Log is our log object.
type Log struct {
	ID     string
	Type   string
	Time   time.Time
	Keys   []string
	Values []string
}
