package internal

import (
	"bytes"
	"runtime/debug"
	"strconv"
)

func GetGoroutineId() int {
	tid := bytes.Fields(debug.Stack())[1]
	res, _ := strconv.Atoi(string(tid))
	return res
}
