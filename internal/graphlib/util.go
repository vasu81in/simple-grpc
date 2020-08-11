package graphlib

import (
	"fmt"
	"runtime"
	"strings"
)

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	}
	return original[i+1:]
}

// WhereAmI ... Helper to print the call stack
func WhereAmI(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("File:%s Func:%s Line:%d",
		chopPath(file), runtime.FuncForPC(function).Name(), line)
}
