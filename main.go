package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	// init the logs map
	logs = make(map[string][]interface{})
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
		commonLog := commonLogFields{}
		for _, d := range lineData {
			if strings.HasPrefix(d, "msg=") {
				id = getLogID(d)
				commonLog.Msg = d
			}
			if strings.HasPrefix(d, "type=") {
				commonLog.Type = strings.Trim(d, "type=")
			}
		}
		logs[id] = append(logs[id], commonLog)
	}

	countTypes()
}
