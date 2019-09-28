package main

import (
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	// init the logs map
	logs = make(map[string][]log)
	openLog()
}

// this is a temporary function until I figure out some stuff
func openLog() {
	file, err := ioutil.ReadFile("test/audit.log")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for _, l := range lines {
		lineData := strings.Split(l, " ")
		id := ""
		tmpLog := log{}
		skip := false
		for _, d := range lineData {
			if d == "" || d == " " || d == "," {
				skip = true
				continue
			}
			if strings.HasPrefix(d, "msg=") {
				id = getLogID(d)
				tmpLog.Time = time.Now()
			} else if strings.HasPrefix(d, "type=") {
				tmpLog.Type = strings.Trim(d, "type=")
			} else {
				keyValue := strings.Split(d, "=")
				tmpLog.Keys = append(tmpLog.Keys, keyValue[0])
				tmpLog.Values = append(tmpLog.Values, keyValue[1])
			}
		}
		if !skip {
			logs[id] = append(logs[id], tmpLog)
		}

	}

	countTypes()
}
