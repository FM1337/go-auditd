package main

import "time"

type log struct {
	Type   string
	Time   time.Time
	Keys   []string
	Values []string
}
