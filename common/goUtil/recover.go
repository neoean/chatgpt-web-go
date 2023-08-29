package goUtil

import (
	"chatgpt-web-new-go/common/logs"
	"runtime"
)

func Recover(source string) func() {
	return func() {
		if err := recover(); err != nil {
			logs.Error("%v panic: %v", source, err)
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			logs.Error("%v panic %s\n", source, string(buf[:n]))
		}
	}
}

func New(f func()) {
	go func() {
		// recover
		defer func() {
			if err := recover(); err != nil {
				var buf [4096]byte
				n := runtime.Stack(buf[:], false)
				logs.Error("goUtil routiune panic %s\n", string(buf[:n]))
			}
		}()

		// do function
		f()
	}()
}
