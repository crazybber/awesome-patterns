package main

import (
	"fmt"
	"os"
)

type Logger struct {
	out *os.File
}

func (l Logger) Log(s string) {
	out := l.out
	if out == nil {
		out = os.Stderr
	}
	fmt.Fprintf(out, "%s [%d]: %s\n", os.Args[0],
		os.Getpid(), s)
}
func (l *Logger) SetOutput(out *os.File) {
	l.out = out
}

func main() {
	l := Logger{}
	l.Log("test")
}
