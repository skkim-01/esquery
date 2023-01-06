package slogger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var g_lockMutex sync.Mutex

func Write(strFile string, args ...interface{}) {
	strText := _buildString(args)
	_write(strFile, strText)
}

// BuildString : Build Log string from args
func _buildString(args ...interface{}) string {
	var strText string
	localTime := time.Now().Local()

	for i := range args {
		if i > 0 {
			strText += "\t"
		}
		strText += fmt.Sprintf("%v", args[i])
	}
	return localTime.Format("2006-01-02 15:04:05") + "\t" + strText[1:len(strText)-1] + "\r\n"
}

func _write(strFile, strText string) {
	g_lockMutex.Lock()
	defer g_lockMutex.Unlock()
	f, e := os.OpenFile(strFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer f.Close()
	f.WriteString(strText)
}
