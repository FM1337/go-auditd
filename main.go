package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	// init the logs map
	logs = make(map[string][]Log)
	initUI()
	openLog()
	renderElements()
}

// this is a temporary function until I figure out some stuff
func openLog() {
	file, err := os.Open("test/audit.log")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		lastKey := ""
		lineData := strings.Split(line, " ")
		id := ""
		tmpLog := Log{}
		skip := false
		for _, d := range lineData {
			if d == "" || d == " " || d == "," {
				skip = true
				continue
			}
			if strings.HasPrefix(d, "msg=audit(") {
				id = getLogID(d)
				if id != lastKey {
					lastKey = id
					keys = append(keys, id)
				}
				tmpLog.ID = id
				tmpLog.Time = time.Unix(getTimestamp(d), 0)
			} else if strings.HasPrefix(d, "type=") {
				tmpLog.Type = strings.Trim(d, "type=")
			} else {
				keyValue := strings.Split(d, "=")
				tmpLog.Keys = append(tmpLog.Keys, keyValue[0])
				tmpLog.Values = append(tmpLog.Values, keyValue[1])
			}
		}
		if !skip {
			// append to the map
			logs[id] = append(logs[id], tmpLog)
		}

	}
	file.Close()
}
