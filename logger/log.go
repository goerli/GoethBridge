package logger

import (
	"fmt"
	"os"
)

func Info(format string, a ...interface{}) {
	out := fmt.Sprintf(format, a...)
	fmt.Println("\x1b[92minfo:\x1b[0m", out)
}

func Warn(format string, a ...interface{}) {
	out := fmt.Sprintf(format, a...)
	fmt.Println("\x1b[93mwarn:\x1b[0m", out)
}

func Error(format string, a ...interface{}) {
	out := fmt.Sprintf(format, a...)
	fmt.Println("\x1b[91merror:\x1b[0m", out)
}

func FatalError(format string, a ...interface{}) {
	out := fmt.Sprintf(format, a...)
	fmt.Println("\x1b[91merror:\x1b[0m", out)	
	os.Exit(1)
}

func Event(format string, a ...interface{}) {
	out := fmt.Sprintf(format, a...)
	fmt.Println("\x1b[94mevent:\x1b[0m", out)
}